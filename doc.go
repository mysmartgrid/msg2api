// Package msg2api is the websocket API of the mysmartgrid2 project and is subdivided into a User API and a Device API.
//
// The user API is used mainly by the sensor graphing page of the web interface.
// Each instance of the graphing page opens a websocket connection linked to the logged in user.
// This connection is then used to request all currently known sensor metadata for display,
// to retreive and receive recorded sensor values, and to request that devices send realtime updates for sensors which support it.
//
// The device API handles all device actions that involve sensors and sensor values.
// With the device API, sensors can be created and removed, sensor metadata can be changed and sensor values can be sent.
package msg2api
