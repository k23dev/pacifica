package cmd

import "fmt"

func Banner(appName, appVersion string) {

	banner := fmt.Sprintf("   %s v(%s)", appName, appVersion)

	separator()
	fmt.Println(banner)
	separator()
	fmt.Println("")

}

func separator() {

	fmt.Println("************************")

}
