package database

import (
	"errors"
)

var (
	ErrCantFindProduct = errors.New("Can't find the product")
	ErrCantDecodeProducts = errors.New("Can't find the product")
	ErrUseerIdIsNotValid = errors.New("The user is not valid")
	ErrCantUpdateUser = errors.New("Cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("Cannot remove this product from the cart")
	ErrCantGetItem = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem = errors.New("cannot update the purchase")
)

func AddProductToCart(){

}

func RemoveCartItem(){

}

func BuyItemFromCart(){

}

func InstantBuy(){

}