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
	// r.HandleFunc("/logout", h.logout)
	// r.HandleFunc("/resetpassword", h.forgotPassword)

	l := r.NewRoute().Subrouter()
	// l.HandleFunc("/registration", h.signUp).Methods("GET")
	// l.HandleFunc("/registration", h.signUpCheck).Methods("POST")
	// l.HandleFunc("/login", h.login).Methods("GET")
	// l.HandleFunc("/login", h.loginCheck).Methods("POST")
	l.Use(h.loginMiddleware)

	s := r.NewRoute().Subrouter()
	// s.Use(h.authMiddleware)
	s.HandleFunc("/category/create", h.createCategories)
	s.HandleFunc("/category/store", h.storeCategories)
	s.HandleFunc("/category/list", h.listCategories)
	s.HandleFunc("/category/{id:[0-9]+}/edit", h.editCategories)
	s.HandleFunc("/category/{id:[0-9]+}/update", h.updateCategories)
	s.HandleFunc("/category/{id:[0-9]+}/delete", h.deleteCategories)
	s.HandleFunc("/post/create", h.createPost)
	s.HandleFunc("/post/store", h.storePost)
	s.HandleFunc("/post/list", h.listPost)
	s.PathPrefix("/cms/asset/").Handler(http.StripPrefix("/cms/asset/", http.FileServer(http.Dir("./"))))
	// s.HandleFunc("/category/search", h.searchCategory)
	// s.HandleFunc("/book/{id:[0-9]+}/edit", h.editBook)
	// s.HandleFunc("/book/{id:[0-9]+}/update", h.updateBook)
	// s.HandleFunc("/book/{id:[0-9]+}/delete", h.deleteBook)
	// s.HandleFunc("/book/search", h.searchBook)
	// s.HandleFunc("/bookings/{id:[0-9]+}/create", h.createBookings)
	// s.HandleFunc("/bookings/store", h.storeBookings)
	// s.HandleFunc("/mybookings", h.myBookings)
	// s.HandleFunc("/book/{id:[0-9]+}/bookdetails", h.bookDetails)
	

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
		// "templates/category/404.html",
		// "templates/book/edit-book.html",
		// "templates/bookings/create-bookings.html",
		// "templates/bookings/my-bookings.html",
		// "templates/book/single-details.html",
		// "templates/signup.html",
		// "templates/login.html",
		// "templates/reset-password.html",
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