# API

## Artists

- `GET /artists` List Artists
- `POST /artists` Create Artist
- `GET /artists/search` Search
- `GET /artists/:id` Details
- `PATCH /artists/:id` Update
- `GET /artists/:id/albums` With Albums
- `GET /artists/:id/tracks` With Tracks

## Albums

- `GET /albums` List Albums
- `POST /albums` Create Album
- `GET /albums/search` Search
- `GET /albums/:id` Details
- `PATCH /albums/:id` Update
- `GET /albums/:id/tracks` With Tracks

## Tracks

- `GET /tracks` List Tracks
- `POST /tracks` Create Track
- `GET /tracks/search` Search
- `GET /tracks/:id` Details
- `PATCH /tracks/:id` Update

## Scrobbles

- `POST /scrobbles` Create Scrobble

## Users

- `GET /users/:id` Details
- `GET /users/:id/scrobbles` List User Scrobbles
- `GET /users/:id/artists` List User Albums
- `GET /users/:id/albums` List User Artists
- `GET /users/:id/tracks` List User Tracks
- `GET /users/:id/now-playing` User Now Playing
