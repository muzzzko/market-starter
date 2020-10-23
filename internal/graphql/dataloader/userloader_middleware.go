package dataloader
//
//import (
//	"context"
//	"net/http"
//	"time"
//	"market-starter/internal/graphql/graph/model"
//)
//
//const loadersKey = "dataloaders"
//
//type Loaders struct {
//	UserById UserLoader
//}
//
//func UserLoaderMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
//			UserById: UserLoader{
//				maxBatch: 100,
//				wait:     1 * time.Millisecond,
//				fetch: func(keys []string) (users []*model.User, errors []error) {
//					return []*model.User{{Name: "name", ID: "1"}}, nil
//				},
//			},
//		})
//		r = r.WithContext(ctx)
//		next.ServeHTTP(w, r)
//	})
//}
//
//func For(ctx context.Context) *Loaders {
//	return ctx.Value(loadersKey).(*Loaders)
//}
