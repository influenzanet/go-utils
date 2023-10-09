package constants

const (
	USER_ROLE_ADMIN           = "ADMIN"
	USER_ROLE_RESEARCHER      = "RESEARCHER"
	USER_ROLE_PARTICIPANT     = "PARTICIPANT"
	USER_ROLE_SERVICE_ACCOUNT = "SERVICE"
)

const (
	EMAIL_TYPE_REGISTRATION           = "registration"
	EMAIL_TYPE_INVITATION             = "invitation"
	EMAIL_TYPE_VERIFY_EMAIL           = "verify-email"
	EMAIL_TYPE_AUTH_VERIFICATION_CODE = "verification-code"
	EMAIL_TYPE_PASSWORD_RESET         = "password-reset"
	EMAIL_TYPE_PASSWORD_CHANGED       = "password-changed"
	EMAIL_TYPE_ACCOUNT_ID_CHANGED     = "account-id-changed"
	EMAIL_TYPE_WEEKLY                 = "weekly"
	EMAIL_TYPE_STUDY_REMINDER         = "study-reminder"
	EMAIL_TYPE_NEWSLETTER             = "newsletter"
	EMAIL_TYPE_ACCOUNT_DELETED        = "account-deleted"
	EMAIL_TYPE_ACCOUNT_INACTIVITY     = "account-inactivity"
)

const (
	TOKEN_PURPOSE_INVITATION                 = "invitation"
	TOKEN_PURPOSE_PASSWORD_RESET             = "password-reset"
	TOKEN_PURPOSE_CONTACT_VERIFICATION       = "contact-verification"
	TOKEN_PURPOSE_SURVEY_LOGIN               = "survey-login"
	TOKEN_PURPOSE_UNSUBSCRIBE_NEWSLETTER     = "unsubscribe-newsletter"
	TOKEN_PURPOSE_RESTORE_ACCOUNT_ID         = "restore_account_id"
	TOKEN_PURPOSE_INACTIVE_USER_NOTIFICATION = "inactive-user-notification"
)
