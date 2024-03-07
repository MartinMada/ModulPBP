package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	UserType int    `json:"type"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Transaction struct {
	ID        int `json:"id"`
	UserID    int `json:"userId"`
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}
type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}
type ProductResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Product `json:"data"`
}
type ProductsResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}
type TransactionResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Transaction `json:"data"`
}
type TransactionsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Transaction `json:"data"`
}

// Detail address
type Address struct {
	ID     int    `json:"id"`
	Street string `json:"street"`
	UserID int    `json:"userid"`
}
type UserAddresses struct {
	User      User      `json:"user"`
	Addresses []Address `json:"addresses"`
}
type UserAddress struct {
	User    User    `json:"user"`
	Address Address `json:"addresses"`
}
type UserAddressesResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []UserAddress `json:"data"`
}

// Detail product
type ProductName struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	ProdID int    `json:"prodid"`
	UserID int    `json:"userid"`
}
type UserProducts struct {
	User     User          `json:"user"`
	Products []ProductName `json:"products"`
}
type UserProduct struct {
	User    User        `json:"user"`
	Product ProductName `json:"products"`
}
type UserProductsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []UserProduct `json:"data"`
}

// Detail transactions
type DetailTransaction struct {
	ID       int     `json:"id"`
	User     User    `json:"user"`
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}
type UserDetailTransactions struct {
	Transactions []DetailTransaction `json:"transactions"`
}
type UserDetailTransaction struct {
	Transaction DetailTransaction `json:"transactions"`
}
type UserDetailTransactionsResponse struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    []UserDetailTransaction `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
