# blog-app

### Run
    - glide install
    - go run src/main/main.go

## Model
    - User
        ID uint64 			`gorm:"primaryKey;autoIncrement" json:"id"`
        Nickname string		`gorm:"size:20;not null;unique" json:"nickname"`
        Email string		`gorm:"size:50;not null;unique" json:"email"`
        Password string		`gorm:"size:60;not null" json:"password,omitempty"` //the field is omitted from the object if its value is empty,
        CreatedAt time.Time	        `gorm:"DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
        UpdatedAt time.Time	        `gorm:"DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
        Posts []Post                `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE" json:"posts,omitempty"` //the field is omitted from the object if its value is empty,
    - Post
        ID uint64 			`gorm:"primaryKey;autoIncrement" json:"id"`
        Title string 		`gorm:"size:30;not null;unique" json:"title"`
        Content string	        `gorm:"size:255;not null;unique" json:"content"`
        Author User			`gorm:"_" json:"author"` //Skip this field for json encoding
        AuthorID uint64             `gorm:"not null" json:"author_id"`
        CreatedAt time.Time	        `gorm:"default currentTimestamp" json:"created_at"`
        UpdatedAt time.Time	        `gorm:"default currentTimestamp" json:"updated_at"`
## Api
    - User
        Uri: "/users",
		Method: GET,
		Description: get all users
        
        Uri: "/users",
		Method: POST,
		Description: Create user

        Uri: "/users/{id}",
		Method: GET,
		Description: Get user

        Uri: "/users/{id}",
		Method: PUT,
		Description: Update user

        Uri: "/users/{id}",
		Method: DELETE,
		Description: Delete user
    - Post
        Uri: "/posts",
		Method: GET,
		Description: get all posts
        
        Uri: "/posts",
		Method: POST,
		Description: Create post

        Uri: "/posts/{id}",
		Method: GET,
		Description: Get post

        Uri: "/posts/{id}",
		Method: PUT,
		Description: Update post

        Uri: "/posts/{id}",
		Method: DELETE,
		Description: Delete post

    - Login
        Uri: "/login",
		Method: http.MethodPost,
		Description: Login (email, password)