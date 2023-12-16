package request

type CreateTagsRequest struct {
	Name string `validate:"required, min=1, max=255" json:"Name"`
}
