package make_redis_key

import "fmt"

func MakeKeyStoreDetail(storeID int32) string {
	return fmt.Sprintf("store_%d_detail", storeID)
}