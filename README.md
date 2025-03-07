# Inventory System - Go Backend API

---

### **📜 Checklist**

### ✅ **1. Project Setup**

- [x]  Initialize Go module: `go mod init inventory-system`
- [x]  Choose a framework (e.g., `Gin`, `Echo`, `Fiber`) (for now starting with `net/http` )
- [ ]  Set up folder structure:
    - More Elaborated File Structure
        
        ---
        
        ## **📂 Project Folder Structure - Inventory System (Go Backend)**
        
        ```
        inventory-system/
        ├── cmd/
        │   ├── main.go                  # Entry point of the application
        │   ├── migrate.go               # Database migration script
        │   ├── seed.go                  # Script to seed initial data
        │
        ├── config/
        │   ├── config.go                 # Configuration loader (from .env)
        │   ├── app.yaml                  # App configuration (alternative to .env)
        │   ├── database.yaml              # Database config settings
        │   ├── logger.yaml                # Logger config settings
        │
        ├── controllers/
        │   ├── auth_controller.go         # Authentication handlers
        │   ├── product_controller.go      # Product-related handlers
        │   ├── category_controller.go     # Category-related handlers
        │   ├── supplier_controller.go     # Supplier-related handlers
        │   ├── inventory_controller.go    # Inventory stock handlers
        │   ├── order_controller.go        # Order processing handlers
        │   ├── transaction_controller.go  # Stock movement handlers
        │   ├── report_controller.go       # Reports & analytics
        │   ├── log_controller.go          # Audit logs & activity tracking
        │
        ├── database/
        │   ├── db.go                      # Database connection logic
        │   ├── migrations/
        │   │   ├── 20240101_create_users.sql
        │   │   ├── 20240102_create_products.sql
        │   │   ├── 20240103_create_orders.sql
        │   │   ├── ... (more migration scripts)
        │
        ├── middleware/
        │   ├── auth_middleware.go         # JWT authentication & user validation
        │   ├── logging_middleware.go      # Logging middleware
        │   ├── rate_limit_middleware.go   # Rate-limiting to prevent abuse
        │   ├── cors_middleware.go         # CORS handling
        │   ├── error_handler.go           # Centralized error handling middleware
        │
        ├── models/
        │   ├── user.go                    # User model (Admin, Staff)
        │   ├── product.go                 # Product model
        │   ├── category.go                # Category model
        │   ├── supplier.go                # Supplier model
        │   ├── inventory.go               # Inventory tracking
        │   ├── order.go                   # Order model
        │   ├── transaction.go             # Transactions (Incoming & Outgoing)
        │   ├── report.go                  # Sales, Stock, Purchase reports
        │   ├── audit_log.go               # Activity logs
        │
        ├── routes/
        │   ├── router.go                   # Main router setup
        │   ├── auth_routes.go              # Auth-related routes
        │   ├── product_routes.go           # Product-related routes
        │   ├── inventory_routes.go         # Inventory management routes
        │   ├── order_routes.go             # Order handling routes
        │   ├── report_routes.go            # Reporting routes
        │   ├── logs_routes.go              # Audit log routes
        │
        ├── services/
        │   ├── auth_service.go             # Handles authentication logic
        │   ├── product_service.go          # Product business logic
        │   ├── category_service.go         # Category-related business logic
        │   ├── supplier_service.go         # Supplier-related logic
        │   ├── inventory_service.go        # Inventory tracking logic
        │   ├── order_service.go            # Order processing logic
        │   ├── transaction_service.go      # Stock movement logic
        │   ├── report_service.go           # Report generation logic
        │   ├── log_service.go              # Handles activity logs
        │
        ├── utils/
        │   ├── hash.go                     # Password hashing utility
        │   ├── jwt.go                      # JWT token generation & validation
        │   ├── logger.go                   # Application logging
        │   ├── response.go                  # Standard API response structure
        │   ├── validation.go                # Input validation helpers
        │   ├── pagination.go                # Pagination utility for large data
        │
        ├── tests/
        │   ├── auth_test.go                 # Test cases for authentication
        │   ├── product_test.go              # Test cases for product management
        │   ├── order_test.go                # Test cases for order processing
        │   ├── inventory_test.go            # Test cases for inventory tracking
        │   ├── transaction_test.go          # Test cases for stock movement
        │
        ├── docs/
        │   ├── swagger.yaml                 # API documentation
        │   ├── postman_collection.json      # Postman collection for testing
        │
        ├── scripts/
        │   ├── setup.sh                     # Setup script for environment
        │   ├── backup.sh                     # Database backup script
        │   ├── deploy.sh                     # Deployment script
        │
        ├── .env                              # Environment variables
        ├── go.mod                            # Go modules
        ├── go.sum                            # Go dependencies
        ├── Dockerfile                        # Docker configuration
        ├── Makefile                          # Build automation tasks
        ├── README.md                         # Project documentation
        
        ```
        
        ---
        
        ## **🛠 Explanation of Key Folders & Files**
        
        ### 📌 **Core Folders**
        
        - `cmd/` → Entry points like `main.go`, migration & seeding scripts
        - `config/` → Configuration files for database, logger, environment settings
        - `controllers/` → Handles HTTP requests, calls `services`, and returns responses
        - `database/` → Handles database connections & migrations
        - `middleware/` → Contains authentication, logging, and security middleware
        - `models/` → Defines database models & ORM configurations
        - `routes/` → Manages API route definitions
        - `services/` → Implements business logic separate from controllers
        - `utils/` → Helper functions for logging, validation, JWT handling, etc.
        - `tests/` → Contains unit and integration test cases
        - `docs/` → API documentation (Swagger/Postman)
        - `scripts/` → Useful scripts for setup, backup, and deployment
        
        ---
        
        ## **🛠 Technologies Used**
        
        | Component | Technology Choices |
        | --- | --- |
        | Backend | Go (Golang) |
        | Framework | Gin / Echo / Fiber |
        | Database | PostgreSQL / MySQL / SQLite |
        | ORM | GORM |
        | Authentication | JWT (JSON Web Token) |
        | Caching | Redis (optional) |
        | Logging | Zap / Logrus |
        | Deployment | Docker, Kubernetes |
        | Testing | Go testing (`testing` package, `testify`) |
        | Documentation | Swagger, Postman |
        
        ---
        
        ## **🚀 Next Steps**
        
        - [ ]  Implement controllers and route logic
        - [ ]  Secure API with authentication & role-based access
        - [ ]  Test with unit & integration tests
        - [ ]  Deploy using Docker/Kubernetes
        
        ---
        
    
    ```
    inventory-system/
    ├── main.go
    ├── config/
    ├── controllers/
    ├── models/
    ├── routes/
    ├── services/
    ├── middleware/
    ├── database/
    ├── utils/
    ├── docs/
    ├── tests/
    
    ```
    
- [x]  Set up `.env` for configuration variables
- [x]  Connect to the database (PostgreSQL, MySQL, SQLite)

---

### 🗄 **2. Database Schema & Models**

- **📌 Table Structures**
    
    ### **1️⃣ Users Table**
    
    | Column | Data Type | Constraints |
    | --- | --- | --- |
    | id | `INTEGER` | PRIMARY KEY, AUTO_INCREMENT |
    | name | `VARCHAR(100)` | NOT NULL |
    | email | `VARCHAR(255)` | UNIQUE, NOT NULL |
    | password | `TEXT` | NOT NULL |
    | role | `ENUM` | ('admin', 'staff', 'manager') DEFAULT 'staff' |
    | created_at | `TIMESTAMP` | DEFAULT CURRENT_TIMESTAMP |
    | updated_at | `TIMESTAMP` | DEFAULT CURRENT_TIMESTAMP |
    
    ---
    
    ### **2️⃣  Categories Table**
    
    | Column | Data Type | Constraints |
    | --- | --- | --- |
    | id | `INTEGER` | PRIMARY KEY, AUTO_INCREMENT |
    | name | `VARCHAR(100)` | UNIQUE, NOT NULL |
    
    ---
    
    ### **3️⃣ Suppliers Table**
    
    | Column | Data Type | Constraints |
    | --- | --- | --- |
    | id | `INTEGER` | PRIMARY KEY, AUTO_INCREMENT |
    | name | `VARCHAR(255)` | NOT NULL |
    | email | `VARCHAR(255)` | UNIQUE, NOT NULL |
    | phone | `VARCHAR(15)` | NOT NULL |
    | address | `TEXT` |  |
    | created_at | `TIMESTAMP` | DEFAULT CURRENT_TIMESTAMP |
    | updated_at | `TIMESTAMP` | DEFAULT CURRENT_TIMESTAMP |
    
    ---
    
    ### **5️⃣ Products Table**
    
    | Column | Data Type | Constraints |
    | --- | --- | --- |
    | id | `INTEGER` | PRIMARY KEY, AUTO_INCREMENT |
    | name | `VARCHAR(255)` | NOT NULL |
    | description | `TEXT` |  |
    | category_id | `INTEGER` | NOT NULL, FOREIGN KEY (categories) |
    | supplier_id | `INTEGER` | NOT NULL, FOREIGN KEY (suppliers) |
    | price | `DECIMAL(10,2)` | NOT NULL |
    | stock | `INTEGER` | DEFAULT 0 |
    | created_at | `TIMESTAMP` | DEFAULT CURRENT_TIMESTAMP |
    | updated_at | `TIMESTAMP` | DEFAULT CURRENT_TIMESTAMP |
    
    ---
    
    ### **4️⃣ Inventory Table**
    
    | Column | Data Type | Constraints |
    | --- | --- | --- |
    | id | `INTEGER` | PRIMARY KEY, AUTO_INCREMENT |
    | product_id | `INTEGER` | NOT NULL, FOREIGN KEY (products) |
    | quantity | `INTEGER` | NOT NULL |
    | updated_at | `TIMESTAMP` | DEFAULT CURRENT_TIMESTAMP |
    
    ---
    
    ### **6️⃣ Transactions Table**
    
    | Column | Data Type | Constraints |
    | --- | --- | --- |
    | id | `INTEGER` | PRIMARY KEY, AUTO_INCREMENT |
    | product_id | `INTEGER` | NOT NULL, FOREIGN KEY (products) |
    | type | `ENUM` | ('add', 'remove') NOT NULL |
    | quantity | `INTEGER` | NOT NULL |
    | reason | `TEXT` |  |
    | created_by | `INTEGER` | FOREIGN KEY (users) |
    | created_at | `TIMESTAMP` | DEFAULT CURRENT_TIMESTAMP |
    
    ---
    
    ### **7️⃣ Orders Table**
    
    | Column | Data Type | Constraints |
    | --- | --- | --- |
    | id | `INTEGER` | PRIMARY KEY, AUTO_INCREMENT |
    | user_id | `INTEGER` | NOT NULL, FOREIGN KEY (users) |
    | status | `ENUM` | ('pending', 'shipped', 'delivered', 'canceled') DEFAULT 'pending' |
    | total | `DECIMAL(10,2)` | NOT NULL |
    | created_at | `TIMESTAMP` | DEFAULT CURRENT_TIMESTAMP |
    | updated_at | `TIMESTAMP` | DEFAULT CURRENT_TIMESTAMP |
    
    ---
    
    ### **8️⃣ Order Items Table**
    
    | Column | Data Type | Constraints |
    | --- | --- | --- |
    | id | `INTEGER` | PRIMARY KEY, AUTO_INCREMENT |
    | order_id | `INTEGER` | NOT NULL, FOREIGN KEY (orders) |
    | product_id | `INTEGER` | NOT NULL, FOREIGN KEY (products) |
    | quantity | `INTEGER` | NOT NULL |
    | price | `DECIMAL(10,2)` | NOT NULL |
    
    ---
    
    ### **9️⃣ Audit Logs Table**
    
    | Column | Data Type | Constraints |
    | --- | --- | --- |
    | id | `INTEGER` | PRIMARY KEY, AUTO_INCREMENT |
    | user_id | `INTEGER` | FOREIGN KEY (users) |
    | action | `VARCHAR(255)` | NOT NULL |
    | table_name | `VARCHAR(100)` |  |
    | record_id | `INTEGER` |  |
    | created_at | `TIMESTAMP` | DEFAULT CURRENT_TIMESTAMP |
- [x]  Create **Users** model (Admin, Staff)
- [x]  Create **Products** model
- [x]  Create **Categories** model
- [x]  Create **Suppliers** model
- [x]  Create **Orders** model (Incoming & Outgoing)
- [x]  Create **Inventory Transactions** model (Stock Movement, Adjustments)
- [x]  Create **Audit Logs** model (Activity Tracking)
- [x]  Migrate database schema using `gorm` or raw SQL

---

### 🔐 **3. Authentication & Authorization**

- [x]  Implement JWT-based authentication
- [x]  Create **Middleware** for role-based access control (RBAC)
- [x]  Implement **User Registration & Login**
- [x]  Secure endpoints with authentication middleware

---

### 🌐 **4. API Endpoints**

Admin —> Manager

Staff ——> 

### 🟢 **Authentication (`/auth`)**

- [x]  `POST /auth/register` - Register a new user
- [x]  `POST /auth/login` - Login and get JWT token
- [ ]  `POST /auth/logout` - Logout user
- [x]  `GET /auth/me` - Get current user details

### 📂 **Categories (`/categories`)**

- [x]  `POST /categories` - Create a category
- [x]  `GET /categories` - Get all categories
- [x]  `GET /categories/:id` - Get category details
- [x]  `PUT /categories/:id` - Update category
- [x]  `DELETE /categories/:id` - Delete category

### 🏢 **Suppliers (`/suppliers`)**

- [ ]  `POST /suppliers` - Add a supplier
- [ ]  `GET /suppliers` - Get all suppliers
- [ ]  `GET /suppliers/:id` - Get supplier details
- [ ]  `PUT /suppliers/:id` - Update supplier
- [ ]  `DELETE /suppliers/:id` - Delete supplier

### 📦 **Products (`/products`)**

- [ ]  `POST /products` - Add a new product
- [ ]  `GET /products` - Get all products
- [ ]  `GET /products/:id` - Get product details
- [ ]  `PUT /products/:id` - Update product
- [ ]  `DELETE /products/:id` - Delete product

### 📜 **Inventory (`/inventory`)**

- [ ]  `GET /inventory` - Get inventory
- [ ]  `GET /inventory/:product_id` - Get inventory details
- [ ]  `POST /inventory/adjust` - Adjust stock manually

### 📦 **Stock Movement (`/transactions`)**

- [ ]  `POST /transactions/incoming` - Add incoming stock
- [ ]  `POST /transactions/outgoing` - Record outgoing stock
- [ ]  `GET /transactions` - Get stock movement history
- [ ]  `GET /transactions/:id` - Get transaction details

### 🛒 **Orders (`/orders`)**

- [ ]  `POST /orders` - Create an order
- [ ]  `GET /orders` - Get all orders
- [ ]  `GET /orders/:id` - Get order details
- [ ]  `PUT /orders/:id` - Update order status
- [ ]  `DELETE /orders/:id` - Delete order

### 📊 **Reports (`/reports`)**

- [ ]  `GET /reports/stock` - Stock reports
- [ ]  `GET /reports/sales` - Sales reports
- [ ]  `GET /reports/purchases` - Purchase reports

### 🔒 **Audit Logs (`/logs`)**

- [ ]  `GET /logs` - Get all activity logs
- [ ]  `GET /logs/:id` - Get log details

---

### ⚙️ **5. Backend Features Implementation**

- [ ]  Implement **CRUD operations**
- [ ]  Secure API with **JWT Middleware**
- [ ]  Implement **Role-based Access Control (RBAC)**
- [ ]  Implement **Pagination & Filtering**
- [ ]  Add **Rate Limiting**
- [ ]  Implement **WebSockets (optional)** for real-time stock updates

---

### 🧪 **6. Testing & Documentation**

- [ ]  Write **Unit Tests**
- [ ]  Write **Integration Tests**
- [ ]  Use **Postman/Swagger** for API documentation
- [ ]  Implement **Logging & Error Handling**

---

### 🚀 **7. Deployment & CI/CD**

- [ ]  Dockerize the application
- [ ]  Set up **CI/CD pipeline**
- [ ]  Deploy on **VPS, Kubernetes, or Cloud provider**
- [ ]  Set up **Monitoring & Alerts**

---

This **Notion-friendly** checklist will help you keep track of your Inventory System project. Let me know if you need modifications! 🚀