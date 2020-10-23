package dependency

type PasswordService interface {
	GetHash(password string) string
	HashIsEqualToPassword(hash string, password string) bool
}
