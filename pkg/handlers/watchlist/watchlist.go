package watchlist

import (
	"fmt"
	"net/http"

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
	router.HandleFunc("/movie/", util.MakeHTTPHandleFunc(h.addMovie)).Methods(http.MethodPost)
}

func (h *Handler) getMovies(w http.ResponseWriter, r *http.Request) error {

	// Get userid from jwt token
	ms, err := h.storage.GetAllMovies(1)
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, ms)

}
func (h *Handler) addMovie(w http.ResponseWriter, r *http.Request) error {
	// Get userid from jwt token
	var mov models.ReqWatchlistItemMovie
	err := util.ParseJSON(r, &mov)
	if err != nil {
		return err
	}
	fmt.Printf("[Debug] Adding movie %#v", mov)

	ms, err := h.storage.GetAllMovies(1)
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, ms)

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
