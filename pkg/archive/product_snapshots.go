package archive

import (
	"context"
	"fmt"
)

// ProductSnapshotParams represents the query parameters for historical product snapshots.
type ProductSnapshotParams struct {
	// ProductId: Required. ID of the product to fetch snapshots for.
	ProductId int `json:"product_id"`

	// Idx: Optional. Only return snapshots with submission_idx <= idx.
	Idx interface{} `json:"idx,omitempty"`

	// MaxTime: Optional. Unix epoch in seconds.
	// Used only if Idx is not provided.
	MaxTime interface{} `json:"max_time,omitempty"`

	// Limit: Optional. Max number of snapshots to return.
	// Defaults to 100. Max possible value is 500.
	Limit int `json:"limit,omitempty"`
}

// ProductSnapshotResponse is the root container for the product snapshot query results.
type ProductSnapshotResponse struct {
	Products []ProductSnapshot `json:"products"`
	Txs      []SnapshotTx      `json:"txs"`
}

// -----------------------------------------------------------------------------
// Section 1: Snapshot Data Structures
// -----------------------------------------------------------------------------

// ProductSnapshot represents the state of a product at a specific submission index.
type ProductSnapshot struct {
	ProductId     int            `json:"product_id"`
	SubmissionIdx string         `json:"submission_idx"`
	Product       SnapshotDetail `json:"product"`
}

// SnapshotDetail can contain either spot or perp product information.
type SnapshotDetail struct {
	Spot *EventProduct `json:"spot,omitempty"` // Reusing EventProduct structure for detailed fields
	Perp *EventProduct `json:"perp,omitempty"`
}

// -----------------------------------------------------------------------------
// Section 2: Transaction Metadata Structures
// -----------------------------------------------------------------------------

// SnapshotTx represents the L1 transaction associated with a product snapshot.
type SnapshotTx struct {
	Tx            interface{} `json:"tx"` // Often contains "update_price" or similar L1 actions
	SubmissionIdx string      `json:"submission_idx"`
	Timestamp     string      `json:"timestamp"` // Crucial for time-series analysis
}

// -----------------------------------------------------------------------------
// Section 3: API Method
// -----------------------------------------------------------------------------

type productSnapshotReqWrapper struct {
	Products ProductSnapshotParams `json:"products"`
}

// GetProductSnapshots queries historical product snapshots ordered by submission index descending.
// Use the submission_idx to associate a product snapshot with its corresponding transaction.
func (c *ArchiveClient) GetProductSnapshots(ctx context.Context, params ProductSnapshotParams) (*ProductSnapshotResponse, error) {
	if params.ProductId <= 0 {
		return nil, fmt.Errorf("archive: valid product_id is required for product snapshots")
	}

	reqPayload := productSnapshotReqWrapper{
		Products: params,
	}

	var resp ProductSnapshotResponse
	if err := c.post(ctx, c.BaseURL, reqPayload, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
