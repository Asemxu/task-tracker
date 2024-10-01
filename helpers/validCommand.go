package helpers

func ValidCommand(arguments []string, cantArguments int) bool {
	return len(arguments) == cantArguments
}
