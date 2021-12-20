package handler

import (
	// "fmt"
	// "math"
	"fmt"
	"net/http"
	"strconv"

	// "strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/mux"

	// "github.com/gorilla/mux"

	tcb "gunkBlog/gunk/v1/category"
)

type Category struct {
	ID int64
	CategoryName string
	IsCompleted bool
	Errors map[string]string
}

type ListCategory struct {
	Categories []Category
	Offset	int
	Limit	int
	Total	int
	Paginate	[]CategoryPagination
	CurrentPage	int
	NextPageURL string
	PreviousPageURL	string
}

type CategoryPagination struct {
	URL string
	PageNumber	int
}

func (c *Category) Validate() error {
	return validation.ValidateStruct(c, validation.Field(
		&c.CategoryName, validation.Required.Error("This field is must be required"),
		validation.Length(3,0).Error("This field is must be grater than 3"),
		))
}

func (h *Handler) createCategories(rw http.ResponseWriter, r *http.Request) {
	vErrs := map[string]string{}
	cat := Category{}
	h.loadCreateCategoryForm(rw, cat.ID, cat.CategoryName, vErrs)
}

func (h *Handler) storeCategories(rw http.ResponseWriter, r *http.Request) {
	if err:= r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	var category Category
	if err := h.decoder.Decode(&category, r.PostForm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := category.Validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[key] = value.Error()
			}
			h.loadCreateCategoryForm(rw, category.ID, category.CategoryName, vErrs)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	_, err := h.tc.Create(r.Context(), &tcb.CreateCategoryRequest{
		Category: &tcb.Category{
			CategoryName: category.CategoryName,
		},
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(rw, r, "/category/list", http.StatusTemporaryRedirect)
}

func (h *Handler) listCategories(rw http.ResponseWriter, r *http.Request) {
	res, err := h.tc.Show(r.Context(), &tcb.ShowCategoryRequest{})
	if err != nil{
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.templates.ExecuteTemplate(rw, "list-category.html", res); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) editCategories(rw http.ResponseWriter, r *http.Request) {
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
	res, err := h.tc.Get(r.Context(), &tcb.GetCategoryRequest{
		ID: Id,
	})
	if err != nil {
		http.Error(rw, "lol", http.StatusInternalServerError)
		return
	}
	rErrs := map[string]string{}
	h.loadEditCategoryForm(rw, res.Category.GetID(), res.Category.CategoryName, rErrs)
}

func (h *Handler) updateCategories(rw http.ResponseWriter, r *http.Request) {
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

	if err := r.ParseForm(); err != nil {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}
	var cat Category
	if err := h.decoder.Decode(&cat, r.PostForm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := cat.Validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[key] = value.Error()
			}
			h.loadEditCategoryForm(rw, cat.ID, cat.CategoryName, vErrs)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	_, err = h.tc.Update(r.Context(), &tcb.UpdateCategoryRequest{
		Category: &tcb.Category{
			ID:           Id,
			CategoryName: cat.CategoryName,
			IsCompleted:  cat.IsCompleted,
		},
	})
	// fmt.Println(Category)
	
	if err != nil{
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(rw, r, "/category/list", http.StatusTemporaryRedirect)
}

func (h *Handler) deleteCategories(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(rw, "lol", http.StatusInternalServerError)
		return
	}

	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("######: %v", id)
	_, err = h.tc.Delete(r.Context(), &tcb.DeleteCategoryRequest{
		ID: Id,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(rw, r, "/category/list", http.StatusTemporaryRedirect)
}

func (h *Handler) loadCreateCategoryForm(rw http.ResponseWriter, id int64, CategoryName string, errs map[string]string) {
	form := Category{
		ID : id,
		CategoryName: CategoryName,
		Errors : errs,
	}
	if err:= h.templates.ExecuteTemplate(rw, "create-category.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) loadEditCategoryForm(rw http.ResponseWriter, id int64, CategoryName string, errs map[string]string) {
	form := Category{
		ID : id,
		CategoryName: CategoryName,
		Errors : errs,
	}

	if err:= h.templates.ExecuteTemplate(rw, "edit-category.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
