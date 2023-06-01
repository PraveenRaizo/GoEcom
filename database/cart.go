package controllers

import "errors"

var (
	ErrCantFindProduct    = errors.New("Can't find the product")
	ErrCantDecodeProducts = errors.New("Can't find the product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("was unable to remove the item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func AddProductToCart() {

}

func RemoveItemFromCart() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
