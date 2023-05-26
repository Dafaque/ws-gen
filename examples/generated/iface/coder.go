// Code generated by wsgen. DO NOT EDIT.
package iface
import "errors"
//FIXME Encoding json not supported
type unimplementedEncoder struct{}
func (ue unimplementedEncoder) Unmarshal([]byte, interface{}) error {
	return errors.New("Unmarshal() is not emplemented")
}
func (ue unimplementedEncoder) Marshal(interface{}) ([]byte, error) {
	return nil, errors.New("Marshal() is not emplemented")
}
var encoder unimplementedEncoder
type Coder interface {
    Unmarshal([]byte, interface{}) error
    Marshal(interface{}) ([]byte, error)
}

type DefaultCoder struct {}
func (dc DefaultCoder) Unmarshal(d []byte, v interface{}) error {
    return encoder.Unmarshal(d, v)
}
func (dc DefaultCoder) Marshal(v interface{}) ([]byte, error) {
    return encoder.Marshal(v)
}