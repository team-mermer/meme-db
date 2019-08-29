# meme-db
meme database backend

## Usage
```bash
cd $WORKDIR
go build main.go
./main $CONFIG_PATH
```

## API Format
* /get_meme_details
    * request
    ```json
    {
        "meme_ids": [
            0,
            1,
            2
        ],
        "n_result": 100
    }
    ```
    * response
    ```json
    [
        {
            "id": 0,
            "title": "chuan",
            "image_url": "....",
            "about": "ark chuan",
            "tags": [
                "ark",
                "chuan"
            ]
        },
        ...
    ]
    ```

* /get_trending_memes
    * request
    ```json
    {
        "n_result": 100
    }
    ```
    * response
    ```json
    [
        {
            "id": 0,
            "title": "chuan",
            "image_url": "....",
            "about": "ark chuan",
            "tags": [
                "ark",
                "chuan"
            ]
        },
        ...
    ]
    ```

* /get_meme_without_tags
    * request
    ```json
    {
        "meme_ids": [
            0,
            1,
            2
        ],
        "n_result": 100
    }
    ```
    * response
    ```json
    [
        {
            "id": 0,
            "title": "chuan",
            "image_url": "....",
            "about": "ark chuan",
        },
        ...
    ]
    ```

* /insert_meme_without_tags
    * request
    ```json
    [
        {
            "id": 0,
            "title": "chuan",
            "image_url": "....",
            "about": "ark chuan",
        },
        ...
    ]
    ```

* /insert_meme_abouts_and_tags
    * request
    ```json
    [
        {
            "id": 0,
            "title": "chuan",
            "image_url": "....",
            "about": "ark chuan",
            "tags": [
                "ark",
                "chuan"
            ]
        },
        ...
    ]
    ```

* /increment_meme_click
    * request
    ```json
    {
        "meme_ids": [
            0,
            1,
            2
        ],
    }
    ```