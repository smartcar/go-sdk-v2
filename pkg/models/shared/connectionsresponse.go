// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConnectionsResponse struct {
	Connections []GetConnection `json:"connections,omitempty"`
	Paging      *CursorPaging   `json:"paging,omitempty"`
}

func (o *ConnectionsResponse) GetConnections() []GetConnection {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *ConnectionsResponse) GetPaging() *CursorPaging {
	if o == nil {
		return nil
	}
	return o.Paging
}
