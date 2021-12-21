package handler

import (
	
	// "log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"

	tcb "gunkBlog/gunk/v1/category"
	tpb "gunkBlog/gunk/v1/post"
)

const sessionName = "cms-session"

type Handler struct {
	templates *template.Template
	decoder *schema.Decoder
	sess *sessions.CookieStore
	tc	tcb.CategoryServiceClient
	tp	tpb.PostServiceClient
}

func New(decoder *schema.Decoder, sess *sessions.CookieStore, tc tcb.CategoryServiceClient, tp tpb.PostServiceClient) *mux.Router {
	h:= &Handler{
		decoder: decoder,
		sess: sess,
		tc: tc,
		tp: tp,
	}

	h.parseTemplate()

	r:= mux.NewRouter()
	r.HandleFunc("/", h.home)

	l := r.NewRoute().Subrouter()
	l.Use(h.loginMiddleware)

	s := r.NewRoute().Subrouter()
	s.HandleFunc("/category/create", h.createCategories)
	s.HandleFunc("/category/store", h.storeCategories)
	s.HandleFunc("/category/list", h.listCategories)
	s.HandleFunc("/category/{id:[0-9]+}/edit", h.editCategories)
	s.HandleFunc("/category/{id:[0-9]+}/update", h.updateCategories)
	s.HandleFunc("/category/{id:[0-9]+}/delete", h.deleteCategories)
	s.HandleFunc("/post/create", h.createPost)
	s.HandleFunc("/post/store", h.storePost)
	s.HandleFunc("/post/list", h.listPost)
	s.HandleFunc("/post/{id:[0-9]+}/edit", h.editPost)
	s.PathPrefix("/cms/asset/").Handler(http.StripPrefix("/cms/asset/", http.FileServer(http.Dir("./"))))
	s.HandleFunc("/post/{id:[0-9]+}/update", h.updatePost)
	s.HandleFunc("/post/{id:[0-9]+}/delete", h.deletePost)
	s.HandleFunc("/post/{id:[0-9]+}/details", h.postDetails)
	

	r.NotFoundHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := h.templates.ExecuteTemplate(rw, "404.html", nil); err != nil {
			http.Error(rw, "invalid URL", http.StatusInternalServerError)
			return
		}
	})

	return r
}

func (h *Handler) parseTemplate() {
	h.templates = template.Must(template.ParseFiles(
		"cms/assets/templates/category/create-category.html",
		"cms/assets/templates/home.html",
		"cms/assets/templates/category/list-category.html",
		"cms/assets/templates/category/edit-category.html",
		"cms/assets/templates/post/create-post.html",
		"cms/assets/templates/post/list-post.html",
		"cms/assets/templates/post/edit-post.html",
		"cms/assets/templates/post/single-post.html",
		))
}

// func (h *Handler) authMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 		session, err := h.sess.Get(r, sessionName)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		authUserID := session.Values["authUserID"]
// 		if authUserID != nil {
// 			next.ServeHTTP(rw, r)
// 		} else {
// 			http.Redirect(rw, r, "/login", http.StatusTemporaryRedirect)
// 		}
		
// 	})
// }

func (h *Handler) loginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(rw, r)
		return
		// session, err := h.sess.Get(r, sessionName)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// authUserID := session.Values["authUserID"]
		// if authUserID != nil {
		// 	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
		// 	return
		// } else {
		// 	next.ServeHTTP(rw, r)
		// }
	})
}