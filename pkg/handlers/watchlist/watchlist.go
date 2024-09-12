package watchlist

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xSaCh/dweep/pkg/models"
	"github.com/xSaCh/dweep/pkg/storage"
	"github.com/xSaCh/dweep/util"
)

type Handler struct {
	storage storage.Storage
}

func NewHandler(storage storage.Storage) *Handler { return &Handler{storage: storage} }

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/movies/", util.MakeHTTPHandleFunc(h.getMovies)).Methods(http.MethodGet)
	router.HandleFunc("/movie/{id}", util.MakeHTTPHandleFunc(h.getMovie)).Methods(http.MethodGet)
	router.HandleFunc("/movie/{id}", util.MakeHTTPHandleFunc(h.addMovie)).Methods(http.MethodPost)
	router.HandleFunc("/movie/{id}", util.MakeHTTPHandleFunc(h.updateMovie)).Methods(http.MethodPut)
	router.HandleFunc("/movie/{id}", util.MakeHTTPHandleFunc(h.deleteMovie)).Methods(http.MethodDelete)
}

func (h *Handler) getMovies(w http.ResponseWriter, r *http.Request) error {

	// Get userid from jwt token
	ms, err := h.storage.WLGetAllMovies(1)
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, ms)

}

func (h *Handler) getMovie(w http.ResponseWriter, r *http.Request) error {
	// Get userid from jwt token
	vars := mux.Vars(r)
	str, ok := vars["id"]
	if !ok {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "missing movie ID"})
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid movie ID"})
	}

	m, err := h.storage.WLGetMovie(id, 1)
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, m)
}

func (h *Handler) addMovie(w http.ResponseWriter, r *http.Request) error {
	// Get userid from jwt token

	vars := mux.Vars(r)
	str, ok := vars["id"]
	if !ok {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "missing movie ID"})
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid movie ID"})
	}

	var mov models.ReqWatchlistItemMovie
	err = util.ParseJSON(r, &mov)
	if err != nil {
		return err
	}
	fmt.Printf("[Debug] Adding movie %#v", mov)

	if err := h.storage.WLAddMovie(mov, id, 1); err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, map[string]string{"msg": "movie added"})

}

func (h *Handler) updateMovie(w http.ResponseWriter, r *http.Request) error {
	// Get userid from jwt token
	vars := mux.Vars(r)
	str, ok := vars["id"]
	if !ok {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "missing movie ID"})
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid movie ID"})
	}

	var mov models.ReqWatchlistItemMovie
	err = util.ParseJSON(r, &mov)
	if err != nil {
		return err
	}
	fmt.Printf("[Debug] Updating movie %#v", mov)

	if err := h.storage.WLUpdateMovie(mov, id, 1); err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, map[string]string{"msg": "movie updated"})
}

func (h *Handler) deleteMovie(w http.ResponseWriter, r *http.Request) error {
	// Get userid from jwt token
	vars := mux.Vars(r)
	str, ok := vars["id"]
	if !ok {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "missing movie ID"})
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid movie ID"})
	}

	if err := h.storage.WLRemoveMovie(id, 1); err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, map[string]string{"msg": "movie removed"})
}

// func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
// 	userID := auth.GetUserIDFromContext(r.Context())

// 	var cart types.CartCheckoutPayload
// 	if err := utils.ParseJSON(r, &cart); err != nil {
// 		utils.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	if err := utils.Validate.Struct(cart); err != nil {
// 		errors := err.(validator.ValidationErrors)
// 		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
// 		return
// 	}

// 	productIds, err := getCartItemsIDs(cart.Items)
// 	if err != nil {
// 		utils.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	// get products
// 	products, err := h.store.GetProductsByID(productIds)
// 	if err != nil {
// 		utils.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	orderID, totalPrice, err := h.createOrder(products, cart.Items, userID)
// 	if err != nil {
// 		utils.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
// 		"total_price": totalPrice,
// 		"order_id":    orderID,
// 	})
// }
