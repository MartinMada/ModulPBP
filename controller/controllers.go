package controller

import (
	m "Modul4_Tugas/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users"
	name := r.URL.Query()["name"]
	age := r.URL.Query()["age"]

	if name != nil {
		fmt.Println(name[0])
		query += " WHERE name='" + name[0] + "'"

	}
	if age != nil {
		if name[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += "age='" + age[0] + "'"
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	var user m.User
	var users []m.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.UserType, &user.Password, &user.Email); err != nil {
			log.Println(err)
			return
		} else {
			users = append(users, user)
		}
	}
	w.Header().Set("Content-Type", "application/json")

	var response m.UsersResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = users
	json.NewEncoder(w).Encode(response)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM products"
	name := r.URL.Query()["name"]
	price := r.URL.Query()["price"]

	if name != nil {
		fmt.Println(name[0])
		query += " WHERE name='" + name[0] + "'"

	}
	if price != nil {
		if name[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += "age='" + price[0] + "'"
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	var product m.Product
	var products []m.Product
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			log.Println(err)
			return
		} else {
			products = append(products, product)
		}
	}
	w.Header().Set("Content-Type", "application/json")

	var response m.ProductsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = products
	json.NewEncoder(w).Encode(response)
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM transactions"
	userID := r.URL.Query()["userID"]
	productID := r.URL.Query()["productID"]

	if userID != nil {
		fmt.Println(userID[0])
		query += " WHERE name='" + userID[0] + "'"

	}
	if productID != nil {
		if userID[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += "age='" + productID[0] + "'"
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	var transaction m.Transaction
	var transactions []m.Transaction
	for rows.Next() {
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Quantity); err != nil {
			log.Println(err)
			return
		} else {
			transactions = append(transactions, transaction)
		}
	}
	w.Header().Set("Content-Type", "application/json")

	var response m.TransactionsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = transactions
	json.NewEncoder(w).Encode(response)
}

// Insert
func InsertUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")

	_, errQuery := db.Exec("INSERT INTO users(name,age,address) values (?,?,?)",
		name,
		age,
		address,
	)

	var response m.UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 400
		response.Message = "Insert Failed"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	name := r.Form.Get("name")
	price, _ := strconv.Atoi(r.Form.Get("price"))

	_, errQuery := db.Exec("INSERT INTO products(name,price) values (?,?)",
		name,
		price,
	)

	var response m.ProductResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 400
		response.Message = "Insert Failed"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// name := r.Form.Get("name")
	userID, _ := strconv.Atoi(r.Form.Get("userid"))
	productID, _ := strconv.Atoi(r.Form.Get("productid"))
	quantity, _ := strconv.Atoi(r.Form.Get("quantity"))

	// address := r.Form.Get("address")

	_, errQuery := db.Exec("INSERT INTO transactions(userid,productid,quantity) values (?,?,?)",
		userID,
		productID,
		quantity,
	)

	var response m.TransactionResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		// response.Status = 400
		// response.Message = "Insert Failed"
		name := r.Form.Get("name")
		price, _ := strconv.Atoi(r.Form.Get("price"))
		_, errQuery := db.Exec("INSERT INTO products(name,price) values (?,?)",
			name,
			price,
		)
		if errQuery == nil {
			response.Status = 200
			response.Message = "Success"
			_, errQuery := db.Exec("INSERT INTO transactions(userid,productid,quantity) values (?,?,?)",
				userID,
				productID,
				quantity,
			)
			if errQuery == nil {
				response.Status = 200
				response.Message = "Success"
			} else {
				response.Status = 400
				response.Message = "Insert Failed"
			}
		} else {
			response.Status = 400
			response.Message = "Insert Failed"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Delete
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// vars := mux.Vars(r)
	// userId := vars["user_id"]
	userId, _ := strconv.Atoi(r.Form.Get("userid"))

	_, errQuery := db.Exec("DELETE FROM users WHERE ID=?",
		userId,
	)

	if errQuery == nil {
		sendSuccessResponseUser(w)
	} else {
		sendErrorResponseUser(w)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	// vars := mux.Vars(r)
	// productId := vars["product_id"]
	productId, _ := strconv.Atoi(r.Form.Get("productid"))

	_, errQuery := db.Exec("DELETE FROM products WHERE id=?",
		productId,
	)

	if errQuery == nil {
		sendSuccessResponseProduct(w)
	} else {
		transactionId, _ := strconv.Atoi(r.Form.Get("transactionid"))

		_, errQuery := db.Exec("DELETE FROM transactions WHERE id=?",
			transactionId,
		)

		if errQuery == nil {
			sendSuccessResponseTransaction(w)
			_, errQuery := db.Exec("DELETE FROM products WHERE id=?",
				productId,
			)
			if errQuery == nil {
				sendSuccessResponseProduct(w)
			} else {
				sendErrorResponseProduct(w)
			}
		} else {
			sendErrorResponseTransaction(w)
		}
	}
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// vars := mux.Vars(r)
	// transactionId := vars["transaction_id"]
	transactionId, _ := strconv.Atoi(r.Form.Get("transactionid"))

	_, errQuery := db.Exec("DELETE FROM transactions WHERE id=?",
		transactionId,
	)

	if errQuery == nil {
		sendSuccessResponseTransaction(w)
	} else {
		sendErrorResponseTransaction(w)
	}
}

// Update
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// vars := mux.Vars(r)
	// userId := vars["user_id"]
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")
	userId, _ := strconv.Atoi(r.Form.Get("userid"))

	// Update data pengguna di database
	_, errQuery := db.Exec("UPDATE users SET name=?, age=?, address=? WHERE id=?",
		name,
		age,
		address,
		userId,
	)

	if errQuery == nil {
		sendSuccessResponseUser(w)
	} else {
		sendErrorResponseUser(w)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	name := r.Form.Get("name")
	price, _ := strconv.Atoi(r.Form.Get("price"))
	productId, _ := strconv.Atoi(r.Form.Get("productid"))

	// Update data pengguna di database
	_, errQuery := db.Exec("UPDATE products SET name=?, price=? WHERE id=?",
		name,
		price,
		productId,
	)

	if errQuery == nil {
		sendSuccessResponseProduct(w)
	} else {
		sendErrorResponseProduct(w)
	}
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	userID, _ := strconv.Atoi(r.Form.Get("userid"))
	productID, _ := strconv.Atoi(r.Form.Get("productid"))
	quantity, _ := strconv.Atoi(r.Form.Get("quantity"))
	transactionId, _ := strconv.Atoi(r.Form.Get("transactionid"))

	// Update data pengguna di database
	_, errQuery := db.Exec("UPDATE transactions SET userid=?, productid=?, quantity=? WHERE id=?",
		userID,
		productID,
		quantity,
		transactionId,
	)

	if errQuery == nil {
		sendSuccessResponseTransaction(w)
	} else {
		sendErrorResponseTransaction(w)
	}
}

// Get User Addresses
func GetUserAddresses(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `SELECT u.ID,u.Name,u.Age,u.Address,u.type,da.ID,da.Street,da.UserID FROM users u
	JOIN detail_address da ON u.ID = da.UserID`
	userAddressRow, err := db.Query(query)
	if err != nil {
		print(err.Error())
		sendErrorResponse(w)
		return
	}

	var userAddress m.UserAddress
	var userAddresses []m.UserAddress
	for userAddressRow.Next() {
		if err := userAddressRow.Scan(
			&userAddress.User.ID, &userAddress.User.Name, &userAddress.User.Age,
			&userAddress.User.Address, &userAddress.User.UserType, &userAddress.Address.ID,
			&userAddress.Address.Street, &userAddress.Address.UserID); err != nil {
			print(err.Error())
			sendErrorResponse(w)
			return
		} else {
			userAddresses = append(userAddresses, userAddress)
		}
	}
	var response m.UserAddressesResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = userAddresses
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// Get User Products
func GetUserProducts(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `SELECT u.ID,u.Name,u.Age,u.Address,u.type,dp.ID,dp.Name,dp.ProdID,dp.UserID FROM users u
	JOIN detail_product dp ON u.ID = dp.UserID ORDER BY u.ID ASC`
	userProductRow, err := db.Query(query)
	if err != nil {
		print(err.Error())
		sendErrorResponse(w)
		return
	}

	var userProduct m.UserProduct
	var UserProducts []m.UserProduct
	for userProductRow.Next() {
		if err := userProductRow.Scan(
			&userProduct.User.ID, &userProduct.User.Name, &userProduct.User.Age,
			&userProduct.User.Address, &userProduct.User.UserType, &userProduct.Product.ID,
			&userProduct.Product.Name, &userProduct.Product.ProdID, &userProduct.Product.UserID); err != nil {
			print(err.Error())
			sendErrorResponse(w)
			return
		} else {
			UserProducts = append(UserProducts, userProduct)
		}
	}
	var response m.UserProductsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = UserProducts
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Get Detail Transactions
func GetUserDetailTransactions(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `SELECT t.ID, u.ID, u.Name, u.Age, u.Address, p.ID, p.Name, p.Price, t.Quantity 
	FROM transactions t
	JOIN users u ON t.UserID = u.ID
	JOIN products p ON t.ProductID = p.ID`
	userTransactionRow, err := db.Query(query)
	if err != nil {
		print(err.Error())
		sendErrorResponse(w)
		return
	}

	var userTransaction m.UserDetailTransaction
	var UserTransactions []m.UserDetailTransaction
	var transaction m.DetailTransaction
	// var user m.User
	// var product m.Product
	for userTransactionRow.Next() {
		if err := userTransactionRow.Scan(
			&transaction.ID, &transaction.User.ID, &transaction.User.Name, &transaction.User.Age,
			&transaction.User.Address, &transaction.Product.ID, &transaction.Product.Name, &transaction.Product.Price, &transaction.Quantity); err != nil {
			print(err.Error())
			sendErrorResponse(w)
			return
		}
		// transaction.User = user
		// transaction.Product = product
		userTransaction.Transaction = transaction
		UserTransactions = append(UserTransactions, userTransaction)
	}
	var response m.UserDetailTransactionsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = UserTransactions
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Get detail transaction by ID
func GetUserDetailTransactionByID(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	transactionID := r.URL.Query().Get("transactionid")

	query := `SELECT t.ID, u.ID, u.Name, u.Age, u.Address, p.ID, p.Name, p.Price, t.Quantity 
	FROM transactions t
	JOIN users u ON t.UserID = u.ID
	JOIN products p ON t.ProductID = p.ID
	WHERE t.ID=?`
	userTransactionRow, err := db.Query(query, transactionID)
	if err != nil {
		print(err.Error())
		sendErrorResponse(w)
		return
	}

	var userTransaction m.UserDetailTransaction
	var UserTransactions []m.UserDetailTransaction
	var transaction m.DetailTransaction

	for userTransactionRow.Next() {
		if err := userTransactionRow.Scan(
			&transaction.ID, &transaction.User.ID, &transaction.User.Name, &transaction.User.Age,
			&transaction.User.Address, &transaction.Product.ID, &transaction.Product.Name, &transaction.Product.Price, &transaction.Quantity); err != nil {
			print(err.Error())
			sendErrorResponse(w)
			return
		}

		userTransaction.Transaction = transaction
		UserTransactions = append(UserTransactions, userTransaction)
	}
	var response m.UserDetailTransactionsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = UserTransactions
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	password := r.Form.Get("password")
	email := r.Form.Get("email")

	platform := r.Header.Get("platform")
	successMessage := "Success login"

	var user m.User
	err = db.QueryRow("SELECT id, name, age, address, type FROM users WHERE email=? AND password=?", email, password).Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.UserType)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}
	if platform != "" {
		successMessage += fmt.Sprintf(" from %s", platform)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(successMessage))
}

func sendErrorResponse(w http.ResponseWriter) {
	var response m.ErrorResponse
	response.Status = 400
	response.Message = "Failed"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func sendSuccessResponseUser(w http.ResponseWriter) {
	var response m.UserResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func sendErrorResponseUser(w http.ResponseWriter) {
	var response m.UserResponse
	response.Status = 400
	response.Message = "Failed"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponseProduct(w http.ResponseWriter) {
	var response m.ProductResponse
	response.Status = 200
	response.Message = "Success"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func sendErrorResponseProduct(w http.ResponseWriter) {
	var response m.ProductResponse
	response.Status = 400
	response.Message = "Failed"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponseTransaction(w http.ResponseWriter) {
	var response m.TransactionResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func sendErrorResponseTransaction(w http.ResponseWriter) {
	var response m.TransactionResponse
	response.Status = 400
	response.Message = "Failed"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
