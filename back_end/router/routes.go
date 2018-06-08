package router

func NewRouter() *mux.Router {

	//Create main router
	r := mux.NewRouter().StrictSlash(true)

	/**
	 * Routes
	 */
	/*r.Methods("GET").Path("/").HandlerFunc(HelloWorld)
	r.Methods("POST").Path("/item/add").HandlerFunc(AddItem)
	r.Methods("GET").Path("/item/list").HandlerFunc(ListItem)
	r.Methods("POST").Path("/item/update").HandlerFunc(UpdateItem)
	r.Methods("POST").Path("/item/remove").HandlerFunc(RemoveItem)*/
	return r
}
