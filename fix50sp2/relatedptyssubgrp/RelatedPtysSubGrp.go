package relatedptyssubgrp

//New returns an initialized RelatedPtysSubGrp instance
func New() *RelatedPtysSubGrp {
	var m RelatedPtysSubGrp
	return &m
}

//NoRelatedPartySubIDs is a repeating group in RelatedPtysSubGrp
type NoRelatedPartySubIDs struct {
	//RelatedPartySubID is a non-required field for NoRelatedPartySubIDs.
	RelatedPartySubID *string `fix:"1567"`
	//RelatedPartySubIDType is a non-required field for NoRelatedPartySubIDs.
	RelatedPartySubIDType *int `fix:"1568"`
}

//NewNoRelatedPartySubIDs returns an initialized NoRelatedPartySubIDs instance
func NewNoRelatedPartySubIDs() *NoRelatedPartySubIDs {
	var m NoRelatedPartySubIDs
	return &m
}

func (m *NoRelatedPartySubIDs) SetRelatedPartySubID(v string)  { m.RelatedPartySubID = &v }
func (m *NoRelatedPartySubIDs) SetRelatedPartySubIDType(v int) { m.RelatedPartySubIDType = &v }

//RelatedPtysSubGrp is a fix50sp2 Component
type RelatedPtysSubGrp struct {
	//NoRelatedPartySubIDs is a non-required field for RelatedPtysSubGrp.
	NoRelatedPartySubIDs []NoRelatedPartySubIDs `fix:"1566,omitempty"`
}

func (m *RelatedPtysSubGrp) SetNoRelatedPartySubIDs(v []NoRelatedPartySubIDs) {
	m.NoRelatedPartySubIDs = v
}
