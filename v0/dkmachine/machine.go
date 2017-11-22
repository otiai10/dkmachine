package dkmachine

const bin = "docker-machine"

// Machine ...
type Machine struct {
	*Inspection   `json:",inline"`
	CreateOptions *CreateOptions `json:"create_options"`
}
