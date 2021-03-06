package ast

import (
	"github.com/elliotchance/c2go/util"
)

type FormatAttr struct {
	Address      string
	Position     string
	Implicit     bool
	Inherited    bool
	FunctionName string
	Unknown1     int
	Unknown2     int
	Children     []Node
}

func parseFormatAttr(line string) *FormatAttr {
	groups := groupsFromRegex(
		`<(?P<position>.*)>
		(?P<implicit> Implicit)?
		(?P<inherited> Inherited)?
		 (?P<function>\w+)
		 (?P<unknown1>\d+)
		 (?P<unknown2>\d+)`,
		line,
	)

	return &FormatAttr{
		Address:      groups["address"],
		Position:     groups["position"],
		Implicit:     len(groups["implicit"]) > 0,
		Inherited:    len(groups["inherited"]) > 0,
		FunctionName: groups["function"],
		Unknown1:     util.Atoi(groups["unknown1"]),
		Unknown2:     util.Atoi(groups["unknown2"]),
		Children:     []Node{},
	}
}

func (n *FormatAttr) AddChild(node Node) {
	n.Children = append(n.Children, node)
}
