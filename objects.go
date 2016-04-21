package ibclient

import (
	"encoding/json"
)

type Bool bool

type EA map[string]interface{}

type EADefListValue string

type IBBase struct {
	objectType string `json:"-"`
}

type IBObject interface {
	ObjectType() string
}

func (obj *IBBase) ObjectType() string {
	return obj.objectType
}

type NetworkView struct {
	IBBase `json:"-"`
	Ref    string `json:"_ref,omitempty"`
	Name   string `json:"name,omitempty"`
	Ea     EA     `json:"extattrs,omitempty"`
}

func NewNetworkView(nv NetworkView) *NetworkView {
	res := nv
	res.objectType = "networkview"

	return &res
}

type Network struct {
	IBBase      `json:"-"`
	Ref         string `json:"_ref,omitempty"`
	NetviewName string `json:"network_view,omitempty"`
	Cidr        string `json:"network,omitempty"`
	Ea          EA     `json:"extattrs,omitempty"`
}

func NewNetwork(nw Network) *Network {
	res := nw
	res.objectType = "network"

	return &res
}

type NetworkContainer struct {
	IBBase      `json:"-"`
	Ref         string `json:"_ref,omitempty"`
	NetviewName string `json:"network_view,omitempty"`
	Cidr        string `json:"network,omitempty"`
	Ea          EA     `json:"extattrs,omitempty"`
}

func NewNetworkContainer(nc NetworkContainer) *NetworkContainer {
	res := nc
	res.objectType = "networkcontainer"

	return &res
}

type FixedAddress struct {
	IBBase      `json:"-"`
	Ref         string `json:"_ref,omitempty"`
	NetviewName string `json:"network_view,omitempty"`
	Cidr        string `json:"network,omitempty"`
	IPAddress   string `json:"ipv4addr,omitempty"`
	Mac         string `json:"mac,omitempty"`
	Ea          EA     `json:"extattrs,omitempty"`
}

func NewFixedAddress(fixedAddr FixedAddress) *FixedAddress {
	res := fixedAddr
	res.objectType = "fixedaddress"

	return &res
}

type EADefinition struct {
	IBBase             `json:"-"`
	Ref                string           `json:"_ref,omitempty"`
	Comment            string           `json:"comment,omitempty"`
	Flags              string           `json:"flags,omitempty"`
	ListValues         []EADefListValue `json:"list_values,omitempty"`
	Name               string           `json:"name,omitempty"`
	Type               string           `json:"type,omitempty"`
	AllowedObjectTypes []string         `json:"allowed_object_types,omitempty"`
}

func NewEADefinition(eadef EADefinition) *EADefinition {
	res := eadef
	res.objectType = "extensibleattributedef"

	return &res
}

func (ea EA) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	for k, v := range ea {
		value := make(map[string]interface{})
		value["value"] = v
		m[k] = value
	}

	return json.Marshal(m)
}

func (val EADefListValue) MarshalJSON() ([]byte, error) {
	m := make(map[string]string)
	m["value"] = string(val)

	return json.Marshal(m)
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if b {
		return json.Marshal("True")
	}

	return json.Marshal("False")
}

func (v *EADefListValue) UnmarshalJSON(b []byte) (err error) {
	var m map[string]string
	err = json.Unmarshal(b, &m)
	if err != nil {
		return
	}

	*v = EADefListValue(m["value"])
	return
}
