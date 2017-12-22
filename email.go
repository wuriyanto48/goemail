package goemail

// emailSender interface
type emailSender interface {
	send() error
	setTemplate(templateFile string, data interface{}) error
}
