package obj

type Item struct {
	ID       uint64 `json:"item_id"`
	UserID   uint64 `json:"user_id"`
	Name     string `json:"item_name"`
	Category string `json:"category"`
	Type     string `json:"type"`
}
