package requests

type (
	//UserSaveRequest - UserSaveRequest
	UserSaveRequest struct {
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
	//UserFindRequest - UserFindRequest
	UserFindRequest struct {
		ID string `json:"id"`
	}
)
