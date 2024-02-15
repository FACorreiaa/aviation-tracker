package flights

import (
	"context"
	"github.com/FACorreiaa/Aviation-tracker/controller/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryFlights struct {
	pgpool *pgxpool.Pool
}

func NewFlights(
	pgpool *pgxpool.Pool,

) *RepositoryFlights {
	return &RepositoryFlights{
		pgpool: pgpool,
	}
}

//func (r *RepositoryFlights) getFlightsData(ctx context.Context, query string,
//	args ...interface{}) ([]models.LiveFlights, error) {
//	var lf []models.LiveFlights
//
//	rows, err := r.pgpool.Query(ctx, query, args...)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var f models.LiveFlights
//		err := rows.Scan(
//			&f.ID,
//			&f.Flight.Codeshared,
//			&f.Flight.Icao,
//			&f.Flight.Iata,
//			&f.Flight.Number,
//			&f.Aircraft.AircraftRegistration,
//			&f.Aircraft.AircraftIcao24,
//			&f.Aircraft.AircraftIcao,
//			&f.Aircraft.AircraftIata,
//			&f.Airline.Icao,
//			&f.Airline.Iata,
//			&f.Airline.Name,
//			&f.Aircraft.AircraftIcao,
//			&f.Aircraft.AircraftRegistration,
//			&f.Aircraft.AircraftIcao24,
//			&f.Aircraft.AircraftIata,
//			&f.FlightDate,
//			&f.FlightStatus,
//			&f.Live.LiveDirection,
//			&f.Live.LiveIsGround,
//			&f.Live.LiveLongitude,
//			&f.Live.LiveLatitude,
//			&f.Live.LiveAltitude,
//			&f.Live.LiveSpeedVertical,
//			&f.Live.LiveSpeedHorizontal,
//			&f.Live.LiveUpdated,
//			&f.Arrival.Airport,
//			&f.Arrival.Timezone,
//			&f.Arrival.Iata,
//			&f.Arrival.Icao,
//			&f.Arrival.Terminal,
//			&f.Arrival.Gate,
//			&f.Arrival.Baggage,
//			&f.Arrival.Delay,
//			&f.Arrival.Scheduled,
//			&f.Arrival.Estimated,
//			&f.Arrival.Actual,
//			&f.Arrival.EstimatedRunway,
//			&f.Arrival.ActualRunway,
//			&f.Departure.Airport,
//			&f.Departure.Timezone,
//			&f.Departure.Iata,
//			&f.Departure.Icao,
//			&f.Departure.Terminal,
//			&f.Departure.Gate,
//			&f.Departure.Delay,
//			&f.Departure.Scheduled,
//			&f.Departure.Estimated,
//			&f.Departure.Actual,
//			&f.Departure.EstimatedRunway,
//			&f.Departure.ActualRunway,
//			&f.CreatedAt,
//		)
//		if err != nil {
//			return nil, err
//		}
//		lf = append(lf, f)
//	}
//
//	if err := rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return lf, nil
//}

func (r *RepositoryFlights) getFlightsData(ctx context.Context, query string,
	args ...interface{}) ([]models.LiveFlights, error) {
	var lf []models.LiveFlights

	rows, err := r.pgpool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var f models.LiveFlights
		err := rows.Scan(
			&f.Flight.Number,
			&f.FlightDate,
			&f.FlightStatus,
			&f.Live.LiveUpdated,
			&f.ID,
			&f.Arrival.Actual,
			&f.Arrival.ActualRunway,
			&f.Arrival.Airport,
			&f.Arrival.Baggage,
			&f.Arrival.Delay,
			&f.Arrival.Estimated,
			&f.Arrival.Terminal,
			&f.Arrival.Gate,
			&f.Arrival.Timezone,
			&f.Departure.Scheduled,
			&f.Departure.EstimatedRunway,
			&f.Departure.Timezone,
			&f.Departure.Terminal,
			&f.Departure.Gate,
			&f.Departure.Actual,
			&f.Departure.ActualRunway,
			&f.Departure.Airport,
			&f.Departure.Estimated,
			&f.Departure.Delay,
		)
		if err != nil {
			return nil, err
		}
		lf = append(lf, f)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lf, nil
}

func (r *RepositoryFlights) GetAllFlights(ctx context.Context,
	page, pageSize int) ([]models.LiveFlights, error) {
	query := `select DISTINCT ON (f.flight_number)
							       f.flight_number,
							       f.flight_date,
							       f.flight_status,
							       f.live_updated,
							       f.id,
							       COALESCE(f.arrival_actual, 'Not defined') as arrival_actual,
							       COALESCE(f.arrival_actual_runway, 'Not defined') as arrival_actual,
							       f.arrival_airport,
							        COALESCE(f.arrival_baggage, 'Not defined'),
							        COALESCE(f.arrival_delay, 0),
							        f.arrival_estimated,
							        COALESCE(f.arrival_terminal, 'Not defined'),
							        COALESCE(f.arrival_gate, 'Not defined'),
							        f.arrival_timezone,
							        f.departure_scheduled,
							        COALESCE(f.departure_estimated_runway, 'Not defined'),
							       f.departure_timezone,
							       f.departure_terminal,
							       COALESCE(f.departure_gate, 'Not defined'),
							       COALESCE(f.departure_actual, 'Not defined'),
							       COALESCE(f.departure_actual_runway, 'Not defined'),
							       f.departure_airport,
							       f.departure_estimated,
							       COALESCE(f.departure_delay, 0)
							from flights f
							order by flight_number, flight_date asc
							offset $1 limit $2`

	offset := (page - 1) * pageSize

	return r.getFlightsData(ctx, query, offset, pageSize)
}

func (r *RepositoryFlights) GetAllFlightsSum(ctx context.Context) (int, error) {
	var count int
	row := r.pgpool.QueryRow(ctx, `SELECT DISTINCT ON(flight_number) Count(id)
										FROM   flights
										GROUP BY flight_number`)

	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
