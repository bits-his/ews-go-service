package streams

import (
	"fmt"

	twitterstream "github.com/fallenstedt/twitter-stream"
	"github.com/fallenstedt/twitter-stream/rules"
)

func AddRules() {

	tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret(KEY, SECRET).RequestBearerToken()
	if err != nil {
		panic(err)
	}
	api := twitterstream.NewTwitterStream(tok.AccessToken)
	rules := twitterstream.NewRuleBuilder().
		//AddRule("lang:en -is:retweet -is:quote (#golangjobs OR #gojobs)", "golang jobs").
		//AddRule("lang:en -is:retweet -is:quote (#python OR #javascript)", "tech jobs").
		AddRule("lang:en -is:retweet -is:quote (#accident OR #earthquake OR #disaster OR #outbreak OR #flood OR #fire)", "flood un nasa earth accident earthquake").
		Build()

	res, err := api.Rules.Create(rules, false) // dryRun is set to false.

	if err != nil {
		panic(err)
	}

	if res.Errors != nil && len(res.Errors) > 0 {
		//https://developer.twitter.com/en/support/twitter-api/error-troubleshooting
		panic(fmt.Sprintf("Received an error from twitter: %v", res.Errors))
	}

	fmt.Println("I have created rules.")
	printRules(res.Data)
}

func GetRules() {
	tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret(KEY, SECRET).RequestBearerToken()
	if err != nil {
		panic(err)
	}
	api := twitterstream.NewTwitterStream(tok.AccessToken)
	res, err := api.Rules.Get()

	if err != nil {
		panic(err)
	}

	if res.Errors != nil && len(res.Errors) > 0 {
		//https://developer.twitter.com/en/support/twitter-api/error-troubleshooting
		panic(fmt.Sprintf("Received an error from twitter: %v", res.Errors))
	}

	if len(res.Data) > 0 {
		fmt.Println("I found these rules: ")
		printRules(res.Data)
	} else {
		fmt.Println("I found no rules")
	}

}

func DeleteRules() {
	tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret(KEY, SECRET).RequestBearerToken()
	if err != nil {
		panic(err)
	}
	api := twitterstream.NewTwitterStream(tok.AccessToken)
	fmt.Println(api.Rules.Get())

	// use api.Rules.Get to find the ID number for an existing rule
	res, err := api.Rules.Delete(rules.NewDeleteRulesRequest(1590314412291854340, 1590315921331085315, 1590370262997995527), false)

	if err != nil {
		panic(err)
	}

	if res.Errors != nil && len(res.Errors) > 0 {
		//https://developer.twitter.com/en/support/twitter-api/error-troubleshooting
		panic(fmt.Sprintf("Received an error from twitter: %v", res.Errors))
	}

	fmt.Println("I have deleted rules ")
}

func printRules(data []rules.DataRule) {
	for _, datum := range data {
		fmt.Printf("Id: %v\n", datum.Id)
		fmt.Printf("Tag: %v\n", datum.Tag)
		fmt.Printf("Value: %v\n\n", datum.Value)
	}
}
