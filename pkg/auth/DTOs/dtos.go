package DTOs

type UserAccessDTO struct {
	IdProduto   string           `json:"idProduto"`
	ProductName string           `json:"productName"`
	IdUser      string           `json:"idUser"`
	LevelAccess []LevelAccessDTO `json:"levelAccess"`
}

type LevelAccessDTO struct {
	LevelAccessName string `json:"levelAccessName"`
	IdLevelAccess   string `json:"idLevelAccess"`
}

const (
	LIST   = "LIST"
	CREATE = "CREATE"
	UPADTE = "UPDATE"
	DELETE = "DELETE"
)

type PERMISSIONS string

func (e PERMISSIONS) String() string {
	switch e {
	case LIST:
		return "LIST"
	case CREATE:
		return "CREATE"
	case UPADTE:
		return "UPDATE"
	case DELETE:
		return "DELETE"
	default:
		return "UNKNOWN"
	}
}

type AuthPermission struct {
	Name       string
	Permission PERMISSIONS
}
