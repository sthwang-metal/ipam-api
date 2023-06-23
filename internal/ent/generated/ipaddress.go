// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipaddress"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblock"
	"go.infratographer.com/x/gidx"
)

// Represents an ip address node on the graph.
type IPAddress struct {
	config `json:"-"`
	// ID of the ent.
	// The ID of the IP Address.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The ip address.
	IP string `json:"IP,omitempty"`
	// The ID for the ip block for this ip address.
	BlockID gidx.PrefixedID `json:"block_id,omitempty"`
	// The ID for the node this is assigned to.
	NodeID gidx.PrefixedID `json:"node_id,omitempty"`
	// Owner ID of the node this is assigned to.
	NodeOwnerID gidx.PrefixedID `json:"node_owner_id,omitempty"`
	// Reserve the IP without it being assigned.
	Reserved bool `json:"reserved,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IPAddressQuery when eager-loading is set.
	Edges        IPAddressEdges `json:"edges"`
	selectValues sql.SelectValues
}

// IPAddressEdges holds the relations/edges for other nodes in the graph.
type IPAddressEdges struct {
	// IPBlock holds the value of the ip_block edge.
	IPBlock *IPBlock `json:"ip_block,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// IPBlockOrErr returns the IPBlock value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IPAddressEdges) IPBlockOrErr() (*IPBlock, error) {
	if e.loadedTypes[0] {
		if e.IPBlock == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: ipblock.Label}
		}
		return e.IPBlock, nil
	}
	return nil, &NotLoadedError{edge: "ip_block"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IPAddress) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ipaddress.FieldID, ipaddress.FieldBlockID, ipaddress.FieldNodeID, ipaddress.FieldNodeOwnerID:
			values[i] = new(gidx.PrefixedID)
		case ipaddress.FieldReserved:
			values[i] = new(sql.NullBool)
		case ipaddress.FieldIP:
			values[i] = new(sql.NullString)
		case ipaddress.FieldCreatedAt, ipaddress.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IPAddress fields.
func (ia *IPAddress) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ipaddress.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ia.ID = *value
			}
		case ipaddress.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ia.CreatedAt = value.Time
			}
		case ipaddress.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ia.UpdatedAt = value.Time
			}
		case ipaddress.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field IP", values[i])
			} else if value.Valid {
				ia.IP = value.String
			}
		case ipaddress.FieldBlockID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field block_id", values[i])
			} else if value != nil {
				ia.BlockID = *value
			}
		case ipaddress.FieldNodeID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field node_id", values[i])
			} else if value != nil {
				ia.NodeID = *value
			}
		case ipaddress.FieldNodeOwnerID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field node_owner_id", values[i])
			} else if value != nil {
				ia.NodeOwnerID = *value
			}
		case ipaddress.FieldReserved:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field reserved", values[i])
			} else if value.Valid {
				ia.Reserved = value.Bool
			}
		default:
			ia.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the IPAddress.
// This includes values selected through modifiers, order, etc.
func (ia *IPAddress) Value(name string) (ent.Value, error) {
	return ia.selectValues.Get(name)
}

// QueryIPBlock queries the "ip_block" edge of the IPAddress entity.
func (ia *IPAddress) QueryIPBlock() *IPBlockQuery {
	return NewIPAddressClient(ia.config).QueryIPBlock(ia)
}

// Update returns a builder for updating this IPAddress.
// Note that you need to call IPAddress.Unwrap() before calling this method if this IPAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (ia *IPAddress) Update() *IPAddressUpdateOne {
	return NewIPAddressClient(ia.config).UpdateOne(ia)
}

// Unwrap unwraps the IPAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ia *IPAddress) Unwrap() *IPAddress {
	_tx, ok := ia.config.driver.(*txDriver)
	if !ok {
		panic("generated: IPAddress is not a transactional entity")
	}
	ia.config.driver = _tx.drv
	return ia
}

// String implements the fmt.Stringer.
func (ia *IPAddress) String() string {
	var builder strings.Builder
	builder.WriteString("IPAddress(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ia.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ia.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ia.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("IP=")
	builder.WriteString(ia.IP)
	builder.WriteString(", ")
	builder.WriteString("block_id=")
	builder.WriteString(fmt.Sprintf("%v", ia.BlockID))
	builder.WriteString(", ")
	builder.WriteString("node_id=")
	builder.WriteString(fmt.Sprintf("%v", ia.NodeID))
	builder.WriteString(", ")
	builder.WriteString("node_owner_id=")
	builder.WriteString(fmt.Sprintf("%v", ia.NodeOwnerID))
	builder.WriteString(", ")
	builder.WriteString("reserved=")
	builder.WriteString(fmt.Sprintf("%v", ia.Reserved))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (ia IPAddress) IsEntity() {}

// IPAddresses is a parsable slice of IPAddress.
type IPAddresses []*IPAddress
