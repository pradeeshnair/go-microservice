package responses

type (
	//UserSaveResponse - UserSaveResponse
	UserSaveResponse struct {
		ID string `json:"id"`
	}
	//UserFindResponse - UserFindResponse
	UserFindResponse struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
)
