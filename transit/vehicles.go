package transit

import (
	"fmt"
	"strconv"

	backend "github.com/benkoppe/bear-trak-backend/backend"
	"github.com/benkoppe/bear-trak-backend/transit/external_gtfs"
	"github.com/jamespfennell/gtfs"
)

func GetVehicles(staticUrl string, realtimeUrls external_gtfs.RealtimeUrls) ([]backend.Vehicle, error) {
	staticGtfs := external_gtfs.GetStaticGtfs(staticUrl)

	realtimeGtfs, err := external_gtfs.GetRealtimeGtfs(realtimeUrls)
	if err != nil {
		return nil, fmt.Errorf("failed to load realtime gtfs data: %v", err)
	}

	var vehicles []backend.Vehicle

	for _, vehicle := range realtimeGtfs.Vehicles {
		vehicleId := vehicle.GetID()
		tripId := vehicle.GetTrip().ID

		id, err := strconv.Atoi(vehicleId.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse vehicle ID as integer: %v", err)
		}

		routeId, err := strconv.Atoi(tripId.RouteID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse route ID as integer: %v", err)
		}

		nextStopTime := nextStopTime(vehicle)
		lastStopTime := stopTimeBefore(vehicle, *nextStopTime)

		nextStop := stopMatchingTime(nextStopTime, *staticGtfs)
		lastStop := stopMatchingTime(lastStopTime, *staticGtfs)

		vehicles = append(vehicles, backend.Vehicle{
			ID:            id,
			RouteId:       routeId,
			Name:          vehicleId.Label,
			DirectionId:   convertStaticDirectionId(tripId.DirectionID),
			Heading:       int(*vehicle.Position.Bearing),
			Longitude:     float64(*vehicle.Position.Longitude),
			Latitude:      float64(*vehicle.Position.Latitude),
			NextStop:      nextStop.Name,
			LastStop:      nillableStop(lastStop),
			DisplayStatus: vehicle.CurrentStatus.String(),
			LastUpdated:   *vehicle.Timestamp,
		})
	}

	return vehicles, nil
}

func nextStopTime(vehicle gtfs.Vehicle) *gtfs.StopTimeUpdate {
	stopId := vehicle.StopID

	for _, stopTime := range vehicle.GetTrip().StopTimeUpdates {
		if stopTime.StopID == stopId {
			return &stopTime
		}
	}

	return nil
}

func stopTimeBefore(vehicle gtfs.Vehicle, stopTime gtfs.StopTimeUpdate) *gtfs.StopTimeUpdate {
	targetSequence := *stopTime.StopSequence - 1

	for _, stopTime := range vehicle.GetTrip().StopTimeUpdates {
		if *stopTime.StopSequence == targetSequence {
			return &stopTime
		}
	}

	return nil
}

func stopMatchingTime(stopTime *gtfs.StopTimeUpdate, staticGtfs gtfs.Static) *gtfs.Stop {
	for _, stop := range staticGtfs.Stops {
		if stop.Id == *stopTime.StopID {
			return &stop
		}
	}

	return nil
}

func nillableStop(stop *gtfs.Stop) backend.NilString {
	if stop == nil {
		return backend.NilString{Null: true}
	} else {
		return backend.NewNilString(stop.Name)
	}
}

// func getLastStop(vehicle gtfs.Vehicle, staticGtfs gtfs.Static) string {
//   for _, trip := range staticGtfs.Trips {
//     if trip.ID == tripId {
//       trip.
//
//     }
//   }
// }
