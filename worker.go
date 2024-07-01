// [TODO]: write backgorund worker here, a stand-alone app that always updating the DB with new data about currencies.
// [TIP]: you can write background jobs natively using go routines and channels.

package main

// [TODO]: let's start to spawn the currency pair updating processes here constantly
// for pairs that are described in constants/currencies.go.
// It is important hele to also check the last time the currency pair was requirested by the API to spawn an update process immediately then,
// an signal might come to this worker outside.
