# Stanley

## Details

From the list of shows in the request payload, return the ones with DRM enabled (drm: true) and at least one episode (episodeCount > 0).

The returned JSON should have a response key with an array of shows. Each element should have the following fields from the request:

    image - corresponding to image/showImage from the request payload
    slug
    title

## Error Handling

If we send invalid JSON, You'll need to return a JSON response with HTTP status 400 Bad Request, and with a `error` key containing the string Could not decode request. For example:

`{
    "error": "Could not decode request: JSON parsing failed"
}`
