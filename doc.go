// redashclient is a go wrapper to Redash's REST API

// redashclient provides a go API to set up and manager a remote redash client programmatically
//
// Setting Up the configuration structure
//
// redashclient uses a configuration structure to hold data that changes between deploymnets to build such configuration structure run
//
//  configuration := config.Config{Host: "0.0.0.0:5005", APIKey: "<A user api key>"}
//
// Initilizting the client
//
// redashclient uses a configuration structure to hold data that changes between deploymnets to build such configuration structure run
//
//  redashClient := redashclient.NewClient(&config.Config{Host: redashAddr, APIKey: apiKey})
//
package redashclient
