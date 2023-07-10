

# Space Repetition algorithm based from Anki program.


import datetime
import math

# Define the initial difficulty level and ease factor for each card
difficulty_levels = {}
ease_factors = {}
INITIAL_DIFFICULTY_LEVEL = 0.3
INITIAL_EASE_FACTOR = 2.5

# Define the function to calculate the review interval for a given card
def calculate_review_interval(difficulty_level, ease_factor, last_review_date):
    
    
    # Calculate the interval based on the difficulty level and ease factor
    if difficulty_level == 0:
        return 1
    elif difficulty_level == 1:
        return 6
    else:
        interval = round((calculate_review_interval(difficulty_level - 0.2, ease_factor, last_review_date) * ease_factor))
        
        # Add a random factor to the interval to prevent clustering of reviews
        interval += math.ceil(interval * 0.1 * (random.random() - 0.5))
        return interval

# Define the function to update the difficulty level and ease factor for a given card
def update_difficulty_and_ease_factor(difficulty_level, ease_factor, quality):
    # Adjust the difficulty level based on the quality of the user's response
    difficulty_level = max(0, difficulty_level + (0.1 - (5 - quality) * (0.08 + (5 - quality) * 0.02)))
    
    
    # Adjust the ease factor based on the quality of the user's response
    ease_factor = max(1.3, ease_factor + 0.1 - (5 - quality) * (0.08 + (5 - quality) * 0.02))
    return difficulty_level, ease_factor

# Define the function to schedule the next review for a given card
def schedule_next_review(card_id, quality):
    
    # Get the current difficulty level and ease factor for the card
    difficulty_level = difficulty_levels.get(card_id, INITIAL_DIFFICULTY_LEVEL)
    ease_factor = ease_factors.get(card_id, INITIAL_EASE_FACTOR)
    
    
    # Update the difficulty level and ease factor for the card based on the user's response
    difficulty_level, ease_factor = update_difficulty_and_ease_factor(difficulty_level, ease_factor, quality)
    
    # Calculate the review interval for the card
    last_review_date = datetime.datetime.now()
    review_interval = calculate_review_interval(difficulty_level, ease_factor, last_review_date)
    
    # Store the updated difficulty level and ease factor for the card
    difficulty_levels[card_id] = difficulty_level
    ease_factors[card_id] = ease_factor
    
    # Return the date of the next review for the card
    return last_review_date + datetime.timedelta(days=review_interval)
