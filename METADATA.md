
# Awesome Feeds > Meta Data

How many ways to add

- Author
- Title
- Date 

Let's count ;-)



## Person / People

- **creator**            -- Dublin Core Meta Data
- **publisher**          -- Dublin Core Meta Data
- **author**             -- RSS 2.0, Atom, JSON Feed
- **contributor**        -- Atom
- **managingEditor**     -- RSS 2.0 Channel
- **webMaster**          -- RSS 2.0 Channel


## Dates

- **published**         -- Atom
- **pubDate**           -- RSS 2.0
- **date_published**    -- JSON Feed
- **date**              -- Dublin Core Meta Data
- **updated**           -- Atom
- **date_modified**     -- JSON Feed
- **lastBuildDate**     -- RSS 2.0 Channel

## Title

- **title**             -- Atom / RSS 2.0 / JSON Feed
- **name**


_2nd Level Title_

- **subtitle**          -- Atom  
- **tagline**

## Summary

- **summary**          -- Atom / JSON Feed
- **description**      -- RSS 2.0
- **abstract**    
- **excerpt**


## Content

- **content**          -- Atom (Defaults to Text!), RSS Yahoo! Search (Media) Extension 
- **content type="text|html|xhtml|"**   -- Atom (Defaults to Text!)
- **content_text**     -- JSON Feed
- **content_html**     -- JSON Feed
- **content:encoded**  -- RDF Content Module



## Tags / Categories

- **category**   -- RSS 2.0
- **category term=**  -- Atom
- **tags[]**     -- JSON Feed
- **keywords**

_Scheme_

- **scheme**     -- Atom
- **domain**     -- RSS 2.0


## Link

- **url**      -- JSON Feed
- **link**     -- RSS 2.0
- **link rel="alternate"**   -- Atom 


_More Links_

- **home_page_url**   -- JSON Feed
- **feed_url**        -- JSON Feed
- **link rel="self"**    -- Atom (feed url)
## ID

- **id**      -- Atom, JSON Feed
- **guid**    -- RSS 2.0
- **permalink**


## Attachments

- **attachments**    -- JSON Feed
- **enclosure**      -- RSS 2.0
- **link rel="?"**   -- Atom

add more examples here - why? why not?

``` json
"attachments": [
                {
                    "url": "http://therecord.co/downloads/The-Record-sp1e1-ChrisParrish.m4a",
                    "mime_type": "audio/x-m4a",
                    "size_in_bytes": 89970236,
                    "duration_in_seconds": 6629
                }
            ]
```


## More

- add banner image for item / entry?
- add image / icon for feed / channel?
