package client

type MarshalableObject interface {
	Unmarshal(data []byte) (MarshalableObject, error)
	Marshal() ([]byte, error)
}
