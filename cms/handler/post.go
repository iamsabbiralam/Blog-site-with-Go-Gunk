package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/mux"

	tcb "gunkBlog/gunk/v1/category"
	tpb "gunkBlog/gunk/v1/post"
)

type Post struct {
	ID				int64
	CategoryID		int64
	Title			string
	Description		string
	Image			string
	CategoryName	string
}

type PostForm struct {
	Post 		Post
	Category	[]Category
	Errors		map[string]string
}

func (p *Post) Validate() error {
	return validation.ValidateStruct(p, 
		validation.Field(&p.Title, 
			validation.Required.Error("This field is must be required"),
			validation.Length(3,0).Error("This field is must be grater than 3"),
		),
		validation.Field(&p.Description,
			validation.Required.Error("The description Field is Required"),
		),)
}

func(h *Handler) categories(rw http.ResponseWriter, r *http.Request) []Category {
	cat, err := h.tc.Show(r.Context(), &tcb.ShowCategoryRequest{})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return nil
	}
	var categories []Category
	for _, val := range cat.Category {
		categories = append(categories, Category{
			ID:           val.ID,
			CategoryName: val.CategoryName,
		})
	}
	return categories
}

func (h *Handler) createPost(rw http.ResponseWriter, r *http.Request) {
	cat := h.categories(rw, r)
	post := Post{}
	vErrs := map[string]string{}
	h.loadCreatePostForm(rw, post, cat, vErrs)
}

func (h *Handler) storePost(rw http.ResponseWriter, r *http.Request) {
	cat := h.categories(rw, r)
	
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var post Post
	if err:= h.decoder.Decode(&post, r.PostForm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("Image")

	if file == nil {
		vErrs := map[string]string{"Image" : "The image field is required"}
		h.loadCreatePostForm(rw, post, cat, vErrs)
			return
	}

    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
    }
    defer file.Close()

    tempFile, err := ioutil.TempFile("cms/assets/images", "upload-*.png")
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
    }
    defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
    }
	tempFile.Write(fileBytes)
	
	imageName := tempFile.Name()

	if err := post.Validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[key] = value.Error()
			}
			h.loadCreatePostForm(rw, post, cat, vErrs)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.tp.Create(r.Context(), &tpb.CreatePostRequest{
		Post: &tpb.Post{
			CategoryID:   post.CategoryID,
			Title:        post.Title,
			Description:  post.Description,
			Image:        imageName,
		},
	})
	if err != nil{
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(rw, r, "/post/list", http.StatusTemporaryRedirect)
}

func(h *Handler) listPost(rw http.ResponseWriter, r *http.Request) {
	res, err := h.tp.Show(r.Context(), &tpb.ShowPostRequest{})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.templates.ExecuteTemplate(rw, "list-post.html", res); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) editPost(rw http.ResponseWriter, r *http.Request) {
	cat := h.categories(rw, r)
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := h.tp.Get(r.Context(), &tpb.GetPostRequest{
		ID: Id,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(res.Post)
	post := Post{
		ID: 		  res.Post.ID,
		CategoryID:   res.Post.CategoryID,
		Title:        res.Post.Title,
		Description:  res.Post.Description,
		Image:        res.Post.Image,
	}
	h.loadEditPostForm(rw, post, cat, map[string]string{})
}

func(h *Handler) updatePost(rw http.ResponseWriter, r *http.Request) {
	cat := h.categories(rw, r)
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := h.tp.Get(r.Context(), &tpb.GetPostRequest{
		ID: Id,
	})

	if err != nil {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}

	var post Post

	if err := h.decoder.Decode(&post, r.PostForm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("Image")
    
	var imageName string
	
    if err == nil {
		defer file.Close()
		tempFile, err := ioutil.TempFile("cms/assets/images", "upload-*.png")
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		
		tempFile.Write(fileBytes)
		
		imageName = tempFile.Name()

		if err := os.Remove(res.Post.Image); err != nil {
				http.Error(rw, "Unable to delete image", http.StatusInternalServerError)
				return
			}
	} else {
		imageName = res.Post.Image
	}

	if err := post.Validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[key] = value.Error()
			}
			h.loadEditPostForm(rw, post, cat, vErrs)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.tp.Update(r.Context(), &tpb.UpdatePostRequest{
		Post: &tpb.Post{
			ID: 		  Id,
			CategoryID:   post.CategoryID,
			Title:        post.Title,
			Description:  post.Description,
			Image:        imageName,
		},
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(rw, r, "/post/list", http.StatusTemporaryRedirect)
}

func (h *Handler) deletePost(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(rw, "Invalid URL", http.StatusInternalServerError)
		return
	}

	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := h.tp.Get(r.Context(), &tpb.GetPostRequest{
		ID: Id,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := os.Remove(res.Post.Image); err != nil {
		http.Error(rw, "Unable to delete image", http.StatusInternalServerError)
		return
	}
	_, err = h.tp.Delete(r.Context(), &tpb.DeletePostRequest{
		ID: Id,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(rw, r, "/post/list", http.StatusTemporaryRedirect)
}

func (h *Handler) loadCreatePostForm(rw http.ResponseWriter, post Post, cat []Category, errs map[string]string) {
	form := PostForm{
		Post:     post,
		Category: cat,
		Errors:   errs,
	}
	if err:= h.templates.ExecuteTemplate(rw, "create-post.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) loadEditPostForm(rw http.ResponseWriter, post Post, cat []Category, errs map[string]string) {
	form := PostForm{
		Category : cat,
		Post : post,
		Errors : errs,
	}
	if err:= h.templates.ExecuteTemplate(rw, "edit-post.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) postDetails(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := h.tp.Get(r.Context(), &tpb.GetPostRequest{
		ID: Id,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.templates.ExecuteTemplate(rw, "single-post.html", res); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}