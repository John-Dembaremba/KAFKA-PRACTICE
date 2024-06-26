package schema

type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type KafkaConn struct {
	Address       string
	Topic         string
	Partitions uint8
}
