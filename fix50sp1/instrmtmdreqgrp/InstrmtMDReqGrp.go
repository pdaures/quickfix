package instrmtmdreqgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
)

//New returns an initialized InstrmtMDReqGrp instance
func New(norelatedsym []NoRelatedSym) *InstrmtMDReqGrp {
	var m InstrmtMDReqGrp
	m.SetNoRelatedSym(norelatedsym)
	return &m
}

//NoRelatedSym is a repeating group in InstrmtMDReqGrp
type NoRelatedSym struct {
	//Instrument is a required component for NoRelatedSym.
	instrument.Instrument
	//UndInstrmtGrp is a non-required component for NoRelatedSym.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp is a non-required component for NoRelatedSym.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//QuoteType is a non-required field for NoRelatedSym.
	QuoteType *int `fix:"537"`
	//SettlType is a non-required field for NoRelatedSym.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NoRelatedSym.
	SettlDate *string `fix:"64"`
	//MDEntrySize is a non-required field for NoRelatedSym.
	MDEntrySize *float64 `fix:"271"`
}

//NewNoRelatedSym returns an initialized NoRelatedSym instance
func NewNoRelatedSym(instrument instrument.Instrument) *NoRelatedSym {
	var m NoRelatedSym
	m.SetInstrument(instrument)
	return &m
}

func (m *NoRelatedSym) SetInstrument(v instrument.Instrument)          { m.Instrument = v }
func (m *NoRelatedSym) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp) { m.UndInstrmtGrp = &v }
func (m *NoRelatedSym) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp) { m.InstrmtLegGrp = &v }
func (m *NoRelatedSym) SetCurrency(v string)                           { m.Currency = &v }
func (m *NoRelatedSym) SetQuoteType(v int)                             { m.QuoteType = &v }
func (m *NoRelatedSym) SetSettlType(v string)                          { m.SettlType = &v }
func (m *NoRelatedSym) SetSettlDate(v string)                          { m.SettlDate = &v }
func (m *NoRelatedSym) SetMDEntrySize(v float64)                       { m.MDEntrySize = &v }

//InstrmtMDReqGrp is a fix50sp1 Component
type InstrmtMDReqGrp struct {
	//NoRelatedSym is a required field for InstrmtMDReqGrp.
	NoRelatedSym []NoRelatedSym `fix:"146"`
}

func (m *InstrmtMDReqGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
