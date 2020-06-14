package resolver

import (
	"sync"

	"github.com/mrtrom/go-graphql-example-api/graph/generated"
)

type Resolver struct {
	UserChannels map[string]chan string
	Mutex        sync.Mutex
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Subscription() generated.SubscriptionResolver {
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
