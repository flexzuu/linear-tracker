package linear

import (
	"github.com/hasura/go-graphql-client"
)

const StateInProgressID = "eab605c7-62bd-42e3-b706-062eb15f21c5"

type AssignedIssues struct {
	Viewer struct {
		AssignedIssues struct {
			Nodes []struct {
				Assignee struct {
					Name      graphql.String
					AvatarURL graphql.String `graphql:"avatarUrl"`
				}
				ID         graphql.String `graphql:"id"`
				Identifier graphql.String
				Title      graphql.String
				BranchName graphql.String
				State      struct {
					ID   graphql.String `graphql:"id"`
					Name graphql.String
				}
				URL graphql.String `graphql:"url"`
			}
		}
	}
}
