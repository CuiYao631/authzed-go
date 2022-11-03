package entity

type PermResponse struct {
	Token        string
	Relationship struct {
		ObjectType  string
		ObjectId    string
		Relation    string
		SubjectType string
		SubjectId   string
	}
}
type ResourcesResponse struct {
	Token    string
	ObjectId string
}
type SubjectResponse struct {
	Token     string
	SubjectID string
}
