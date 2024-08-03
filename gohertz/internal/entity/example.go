package entity

type SrvExampleList struct {
	Search *string
	Limit  *int
	Offset *int
}

type SrvExampleDetail struct {
	Id *string
}

type SrvExampleCreate struct {
	Id     *string
	Name   *string
	Detail *string
	Cuont  *float64
}

type SrvExamplePut struct {
	Id     *string
	Name   *string
	Detail *string
	Cuont  *float64
}

type SrvExamplePatch struct {
	Id     *string
	Name   *string
	Detail *string
	Cuont  *float64
}

type SrvExampleDelete struct {
	Id *string
}
