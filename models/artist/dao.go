package artists


/// this might need to represent the DAO for the struct above
type ArtistService struct {
	db sql.Conn
}

func InitArtistService(*db sql.Conn) *ArtistService {
	return &ArtistService{
		db: db
	}
}

func (*s ArtistService) GetallArtists() Artist[] {
	db.query()
	// map to the struct
	// retur
}