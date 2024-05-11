package main

import (
	"p1/constant"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func TestAlgo(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (t *TestSuite) TestIsNotGotShotFirst() {
	expect := true
	actual := isNotRevengeFirst(constant.GotShot)
	t.Equal(expect, actual)
}
func (t *TestSuite) TestIsShotFirst() {
	expect := false
	actual := isNotRevengeFirst(constant.Revenge)
	t.Equal(expect, actual)
}

func (t *TestSuite) TestIsLastShotIsGotShot() {
	expect := false
	actual := isLastShotNotIsGotShot(constant.GotShot)
	t.Equal(expect, actual)
}

func (t *TestSuite) TestIsLastShotIsNotGotShot() {
	expect := true
	actual := isLastShotNotIsGotShot(constant.Revenge)
	t.Equal(expect, actual)
}

func (t *TestSuite) TestEachShotWasRevenge() {
	expect := true
	actual := isEachShotWasRevenge(constant.TC1)
	t.Equal(expect, actual)
}

func (t *TestSuite) TestEachShotWasNotRevenge() {
	expect := false
	actual := isEachShotWasRevenge(constant.TC3)
	t.Equal(expect, actual)
}

func (t *TestSuite) TestGoodBoyTC1() {
	expect := true
	actual := isGoodBoy(constant.TC1)
	t.Equal(expect, actual)
}

func (t *TestSuite) TestBadBoyRevengeFirst() {
	expect := false
	actual := isGoodBoy(constant.TC2)
	t.Equal(expect, actual)
}

func (t *TestSuite) TestBadBoyLastGotShot() {
	expect := false
	actual := isGoodBoy(constant.TC3)
	t.Equal(expect, actual)
}

func (t *TestSuite) TestGoodBoyNoShotHappen() {
	expect := false
	actual := isGoodBoy(constant.TC0)
	t.Equal(expect, actual)
}
func (t *TestSuite) TestBadBoySomeNotRevenge() {
	expect := false
	actual := isGoodBoy(constant.TC7)
	t.Equal(expect, actual)
}