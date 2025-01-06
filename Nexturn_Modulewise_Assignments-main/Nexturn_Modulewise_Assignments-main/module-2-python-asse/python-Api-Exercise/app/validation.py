VALID_GENRES = ["Fiction", "Non-Fiction", "Mystery", "Sci-Fi", "Fantasy"]

def validate_book_data(data):
    required_fields = ['title', 'author', 'published_year', 'genre']
    for field in required_fields:
        if field not in data:
            return f"'{field}' is required."
    if not isinstance(data['published_year'], int) or data['published_year'] <= 0:
        return "'published_year' must be a valid year."
    if data['genre'] not in VALID_GENRES:
        return f"'genre' must be one of {', '.join(VALID_GENRES)}."
    return None