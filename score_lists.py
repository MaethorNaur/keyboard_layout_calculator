#   """Description of how the original score list came to be:"""
#   
#   # Define how important layer-placement is as opposed to flow. 0 = only flow is important. 1 =  only what layer the letter is in is important.
#   layerVsFlow = 0.5  # /0.6
#   
#   # Define the comfort of different Layers. Use numbers between 1 (most comfortable) and 0 (least comfortable).
#   L1_comfort = 1
#   L2_comfort = 0.7
#   L3_comfort = 0.3
#   L4_comfort = 0
#   
#   # Define what placement-combinations have a "good flow"
#   # Put in numbers between 1 (best flow) and 0 (worst flow).
#   # 0 (the middle of this array) is assumed to be the position of the first letter. IT'S ASSUMED TO BE EVEN!!!
#   # ( = where the 'e' is in the current Layout)
#   # +1 is one step clockwise. +2 is two steps clockwise. -1 is one step counterclockwise. -2 is two steps counterclockwise.
#   # Place the score-numbers in a way that reflects how well the second letter follows after the first one.
#   # For explanations, see https://github.com/flide/8VIM/discussions/99#discussioncomment-585774 and the following messages.
#   #                 -7  -6   -5   -4  -3  -2   -1  ~0~ 1   2    3    4   5    6    7
#   flow_evenPos_L1 = [0, 0.3, 0.8, 0.5, 1, 0.9, 0.8, 1, 0, 0.3, 0.8, 0.5, 1, 0.9, 0.8]
#   
#   #                  -7   -6  -5  -4   -3   -2    -1  ~0~  1    2   3   4    5    6     7
#   flow_evenPos_L2 = [0.5, 0.9, 0, 0.5, 0.8, 0.5, 0.95, 1, 0.5, 0.9, 0, 0.5, 0.8, 0.5, 0.95]
#   
#   #                 -7 -6  -5   -4  -3  -2   -1   ~0~  1  2   3    4   5   6    7
#   flow_evenPos_L3 = [1, 1, 0.4, 0.9, 0, 0.3, 0.5, 0.5, 1, 1, 0.4, 0.9, 0, 0.3, 0.5]
#   
#   #                  -7   -6    -5  -4  -3   -2  -1  ~0~   1    2    3    4   5    6   7
#   flow_evenPos_L4 = [0.9, 0.5, 0.95, 1, 0.5, 0.9, 0, 0.3, 0.9, 0.5, 0.95, 1, 0.5, 0.9, 0]
ORIGINAL_SCORE_LIST = [
    [1.000, 0.500, 0.650, 0.900, 0.750, 1.000, 0.950, 0.900, 1.000, 0.500, 0.650, 0.900, 0.750, 1.000, 0.950, 0.900, 1.000, 0.500, 0.650, 0.900, 0.750, 1.000, 0.950, 0.900, 1.000, 0.500, 0.650, 0.900, 0.750, 1.000, 0.950, 0.900],
    [0.500, 1.000, 0.900, 0.950, 1.000, 0.750, 0.900, 0.650, 0.500, 1.000, 0.900, 0.950, 1.000, 0.750, 0.900, 0.650, 0.500, 1.000, 0.900, 0.950, 1.000, 0.750, 0.900, 0.650, 0.500, 1.000, 0.900, 0.950, 1.000, 0.750, 0.900, 0.650],
    [0.950, 0.900, 1.000, 0.500, 0.650, 0.900, 0.750, 1.000, 0.950, 0.900, 1.000, 0.500, 0.650, 0.900, 0.750, 1.000, 0.950, 0.900, 1.000, 0.500, 0.650, 0.900, 0.750, 1.000, 0.950, 0.900, 1.000, 0.500, 0.650, 0.900, 0.750, 1.000],
    [0.900, 0.650, 0.500, 1.000, 0.900, 0.950, 1.000, 0.750, 0.900, 0.650, 0.500, 1.000, 0.900, 0.950, 1.000, 0.750, 0.900, 0.650, 0.500, 1.000, 0.900, 0.950, 1.000, 0.750, 0.900, 0.650, 0.500, 1.000, 0.900, 0.950, 1.000, 0.750],
    [0.750, 1.000, 0.950, 0.900, 1.000, 0.500, 0.650, 0.900, 0.750, 1.000, 0.950, 0.900, 1.000, 0.500, 0.650, 0.900, 0.750, 1.000, 0.950, 0.900, 1.000, 0.500, 0.650, 0.900, 0.750, 1.000, 0.950, 0.900, 1.000, 0.500, 0.650, 0.900],
    [1.000, 0.750, 0.900, 0.650, 0.500, 1.000, 0.900, 0.950, 1.000, 0.750, 0.900, 0.650, 0.500, 1.000, 0.900, 0.950, 1.000, 0.750, 0.900, 0.650, 0.500, 1.000, 0.900, 0.950, 1.000, 0.750, 0.900, 0.650, 0.500, 1.000, 0.900, 0.950],
    [0.650, 0.900, 0.750, 1.000, 0.950, 0.900, 1.000, 0.500, 0.650, 0.900, 0.750, 1.000, 0.950, 0.900, 1.000, 0.500, 0.650, 0.900, 0.750, 1.000, 0.950, 0.900, 1.000, 0.500, 0.650, 0.900, 0.750, 1.000, 0.950, 0.900, 1.000, 0.500],
    [0.900, 0.950, 1.000, 0.750, 0.900, 0.650, 0.500, 1.000, 0.900, 0.950, 1.000, 0.750, 0.900, 0.650, 0.500, 1.000, 0.900, 0.950, 1.000, 0.750, 0.900, 0.650, 0.500, 1.000, 0.900, 0.950, 1.000, 0.750, 0.900, 0.650, 0.500, 1.000],
    [0.850, 0.600, 0.800, 0.350, 0.600, 0.750, 0.600, 0.825, 0.850, 0.600, 0.800, 0.350, 0.600, 0.750, 0.600, 0.825, 0.850, 0.600, 0.800, 0.350, 0.600, 0.750, 0.600, 0.825, 0.850, 0.600, 0.800, 0.350, 0.600, 0.750, 0.600, 0.825],
    [0.600, 0.850, 0.825, 0.600, 0.750, 0.600, 0.350, 0.800, 0.600, 0.850, 0.825, 0.600, 0.750, 0.600, 0.350, 0.800, 0.600, 0.850, 0.825, 0.600, 0.750, 0.600, 0.350, 0.800, 0.600, 0.850, 0.825, 0.600, 0.750, 0.600, 0.350, 0.800],
    [0.600, 0.825, 0.850, 0.600, 0.800, 0.350, 0.600, 0.750, 0.600, 0.825, 0.850, 0.600, 0.800, 0.350, 0.600, 0.750, 0.600, 0.825, 0.850, 0.600, 0.800, 0.350, 0.600, 0.750, 0.600, 0.825, 0.850, 0.600, 0.800, 0.350, 0.600, 0.750],
    [0.350, 0.800, 0.600, 0.850, 0.825, 0.600, 0.750, 0.600, 0.350, 0.800, 0.600, 0.850, 0.825, 0.600, 0.750, 0.600, 0.350, 0.800, 0.600, 0.850, 0.825, 0.600, 0.750, 0.600, 0.350, 0.800, 0.600, 0.850, 0.825, 0.600, 0.750, 0.600],
    [0.600, 0.750, 0.600, 0.825, 0.850, 0.600, 0.800, 0.350, 0.600, 0.750, 0.600, 0.825, 0.850, 0.600, 0.800, 0.350, 0.600, 0.750, 0.600, 0.825, 0.850, 0.600, 0.800, 0.350, 0.600, 0.750, 0.600, 0.825, 0.850, 0.600, 0.800, 0.350],
    [0.750, 0.600, 0.350, 0.800, 0.600, 0.850, 0.825, 0.600, 0.750, 0.600, 0.350, 0.800, 0.600, 0.850, 0.825, 0.600, 0.750, 0.600, 0.350, 0.800, 0.600, 0.850, 0.825, 0.600, 0.750, 0.600, 0.350, 0.800, 0.600, 0.850, 0.825, 0.600],
    [0.800, 0.350, 0.600, 0.750, 0.600, 0.825, 0.850, 0.600, 0.800, 0.350, 0.600, 0.750, 0.600, 0.825, 0.850, 0.600, 0.800, 0.350, 0.600, 0.750, 0.600, 0.825, 0.850, 0.600, 0.800, 0.350, 0.600, 0.750, 0.600, 0.825, 0.850, 0.600],
    [0.825, 0.600, 0.750, 0.600, 0.350, 0.800, 0.600, 0.850, 0.825, 0.600, 0.750, 0.600, 0.350, 0.800, 0.600, 0.850, 0.825, 0.600, 0.750, 0.600, 0.350, 0.800, 0.600, 0.850, 0.825, 0.600, 0.750, 0.600, 0.350, 0.800, 0.600, 0.850],
    [0.400, 0.650, 0.650, 0.350, 0.600, 0.150, 0.300, 0.400, 0.400, 0.650, 0.650, 0.350, 0.600, 0.150, 0.300, 0.400, 0.400, 0.650, 0.650, 0.350, 0.600, 0.150, 0.300, 0.400, 0.400, 0.650, 0.650, 0.350, 0.600, 0.150, 0.300, 0.400],
    [0.650, 0.400, 0.400, 0.300, 0.150, 0.600, 0.350, 0.650, 0.650, 0.400, 0.400, 0.300, 0.150, 0.600, 0.350, 0.650, 0.650, 0.400, 0.400, 0.300, 0.150, 0.600, 0.350, 0.650, 0.650, 0.400, 0.400, 0.300, 0.150, 0.600, 0.350, 0.650],
    [0.300, 0.400, 0.400, 0.650, 0.650, 0.350, 0.600, 0.150, 0.300, 0.400, 0.400, 0.650, 0.650, 0.350, 0.600, 0.150, 0.300, 0.400, 0.400, 0.650, 0.650, 0.350, 0.600, 0.150, 0.300, 0.400, 0.400, 0.650, 0.650, 0.350, 0.600, 0.150],
    [0.350, 0.650, 0.650, 0.400, 0.400, 0.300, 0.150, 0.600, 0.350, 0.650, 0.650, 0.400, 0.400, 0.300, 0.150, 0.600, 0.350, 0.650, 0.650, 0.400, 0.400, 0.300, 0.150, 0.600, 0.350, 0.650, 0.650, 0.400, 0.400, 0.300, 0.150, 0.600],
    [0.600, 0.150, 0.300, 0.400, 0.400, 0.650, 0.650, 0.350, 0.600, 0.150, 0.300, 0.400, 0.400, 0.650, 0.650, 0.350, 0.600, 0.150, 0.300, 0.400, 0.400, 0.650, 0.650, 0.350, 0.600, 0.150, 0.300, 0.400, 0.400, 0.650, 0.650, 0.350],
    [0.150, 0.600, 0.350, 0.650, 0.650, 0.400, 0.400, 0.300, 0.150, 0.600, 0.350, 0.650, 0.650, 0.400, 0.400, 0.300, 0.150, 0.600, 0.350, 0.650, 0.650, 0.400, 0.400, 0.300, 0.150, 0.600, 0.350, 0.650, 0.650, 0.400, 0.400, 0.300],
    [0.650, 0.350, 0.600, 0.150, 0.300, 0.400, 0.400, 0.650, 0.650, 0.350, 0.600, 0.150, 0.300, 0.400, 0.400, 0.650, 0.650, 0.350, 0.600, 0.150, 0.300, 0.400, 0.400, 0.650, 0.650, 0.350, 0.600, 0.150, 0.300, 0.400, 0.400, 0.650],
    [0.400, 0.300, 0.150, 0.600, 0.350, 0.650, 0.650, 0.400, 0.400, 0.300, 0.150, 0.600, 0.350, 0.650, 0.650, 0.400, 0.400, 0.300, 0.150, 0.600, 0.350, 0.650, 0.650, 0.400, 0.400, 0.300, 0.150, 0.600, 0.350, 0.650, 0.650, 0.400],
    [0.150, 0.450, 0.250, 0.475, 0.500, 0.250, 0.450, 0.000, 0.150, 0.450, 0.250, 0.475, 0.500, 0.250, 0.450, 0.000, 0.150, 0.450, 0.250, 0.475, 0.500, 0.250, 0.450, 0.000, 0.150, 0.450, 0.250, 0.475, 0.500, 0.250, 0.450, 0.000],
    [0.450, 0.150, 0.000, 0.450, 0.250, 0.500, 0.475, 0.250, 0.450, 0.150, 0.000, 0.450, 0.250, 0.500, 0.475, 0.250, 0.450, 0.150, 0.000, 0.450, 0.250, 0.500, 0.475, 0.250, 0.450, 0.150, 0.000, 0.450, 0.250, 0.500, 0.475, 0.250],
    [0.450, 0.000, 0.150, 0.450, 0.250, 0.475, 0.500, 0.250, 0.450, 0.000, 0.150, 0.450, 0.250, 0.475, 0.500, 0.250, 0.450, 0.000, 0.150, 0.450, 0.250, 0.475, 0.500, 0.250, 0.450, 0.000, 0.150, 0.450, 0.250, 0.475, 0.500, 0.250],
    [0.475, 0.250, 0.450, 0.150, 0.000, 0.450, 0.250, 0.500, 0.475, 0.250, 0.450, 0.150, 0.000, 0.450, 0.250, 0.500, 0.475, 0.250, 0.450, 0.150, 0.000, 0.450, 0.250, 0.500, 0.475, 0.250, 0.450, 0.150, 0.000, 0.450, 0.250, 0.500],
    [0.500, 0.250, 0.450, 0.000, 0.150, 0.450, 0.250, 0.475, 0.500, 0.250, 0.450, 0.000, 0.150, 0.450, 0.250, 0.475, 0.500, 0.250, 0.450, 0.000, 0.150, 0.450, 0.250, 0.475, 0.500, 0.250, 0.450, 0.000, 0.150, 0.450, 0.250, 0.475],
    [0.250, 0.500, 0.475, 0.250, 0.450, 0.150, 0.000, 0.450, 0.250, 0.500, 0.475, 0.250, 0.450, 0.150, 0.000, 0.450, 0.250, 0.500, 0.475, 0.250, 0.450, 0.150, 0.000, 0.450, 0.250, 0.500, 0.475, 0.250, 0.450, 0.150, 0.000, 0.450],
    [0.250, 0.475, 0.500, 0.250, 0.450, 0.000, 0.150, 0.450, 0.250, 0.475, 0.500, 0.250, 0.450, 0.000, 0.150, 0.450, 0.250, 0.475, 0.500, 0.250, 0.450, 0.000, 0.150, 0.450, 0.250, 0.475, 0.500, 0.250, 0.450, 0.000, 0.150, 0.450],
    [0.000, 0.450, 0.250, 0.500, 0.475, 0.250, 0.450, 0.150, 0.000, 0.450, 0.250, 0.500, 0.475, 0.250, 0.450, 0.150, 0.000, 0.450, 0.250, 0.500, 0.475, 0.250, 0.450, 0.150, 0.000, 0.450, 0.250, 0.500, 0.475, 0.250, 0.450, 0.150],
]

# The new score-list, resulting from this post: https://github.com/flide/8VIM/discussions/260#discussioncomment-2376926
KJOETOM_SCORE_LIST = [
    [0.896, 0.683, 0.766, 0.785, 0.698, 0.863, 1.000, 0.785, 0.758, 0.600, 0.663, 0.677, 0.611, 0.734, 0.831, 0.677, 0.664, 0.539, 0.590, 0.601, 0.549, 0.646, 0.720, 0.601, 0.591, 0.490, 0.531, 0.540, 0.498, 0.576, 0.635, 0.540],
    [0.683, 0.896, 0.785, 1.000, 0.863, 0.698, 0.785, 0.766, 0.600, 0.758, 0.677, 0.831, 0.734, 0.611, 0.677, 0.663, 0.539, 0.664, 0.601, 0.720, 0.646, 0.549, 0.601, 0.590, 0.490, 0.591, 0.540, 0.635, 0.576, 0.498, 0.540, 0.531],
    [1.000, 0.785, 0.896, 0.683, 0.766, 0.785, 0.698, 0.863, 0.831, 0.677, 0.758, 0.600, 0.663, 0.677, 0.611, 0.734, 0.720, 0.601, 0.664, 0.539, 0.590, 0.601, 0.549, 0.646, 0.635, 0.540, 0.591, 0.490, 0.531, 0.540, 0.498, 0.576],
    [0.785, 0.766, 0.683, 0.896, 0.785, 1.000, 0.863, 0.698, 0.677, 0.663, 0.600, 0.758, 0.677, 0.831, 0.734, 0.611, 0.601, 0.590, 0.539, 0.664, 0.601, 0.720, 0.646, 0.549, 0.540, 0.531, 0.490, 0.591, 0.540, 0.635, 0.576, 0.498],
    [0.698, 0.863, 1.000, 0.785, 0.896, 0.683, 0.766, 0.785, 0.611, 0.734, 0.831, 0.677, 0.758, 0.600, 0.663, 0.677, 0.549, 0.646, 0.720, 0.601, 0.664, 0.539, 0.590, 0.601, 0.498, 0.576, 0.635, 0.540, 0.591, 0.490, 0.531, 0.540],
    [0.863, 0.698, 0.785, 0.766, 0.683, 0.896, 0.785, 1.000, 0.734, 0.611, 0.677, 0.663, 0.600, 0.758, 0.677, 0.831, 0.646, 0.549, 0.601, 0.590, 0.539, 0.664, 0.601, 0.720, 0.576, 0.498, 0.540, 0.531, 0.490, 0.591, 0.540, 0.635],
    [0.766, 0.785, 0.698, 0.863, 1.000, 0.785, 0.896, 0.683, 0.663, 0.677, 0.611, 0.734, 0.831, 0.677, 0.758, 0.600, 0.590, 0.601, 0.549, 0.646, 0.720, 0.601, 0.664, 0.539, 0.531, 0.540, 0.498, 0.576, 0.635, 0.540, 0.591, 0.490],
    [0.785, 1.000, 0.863, 0.698, 0.785, 0.766, 0.683, 0.896, 0.677, 0.831, 0.734, 0.611, 0.677, 0.663, 0.600, 0.758, 0.601, 0.720, 0.646, 0.549, 0.601, 0.590, 0.539, 0.664, 0.540, 0.635, 0.576, 0.498, 0.540, 0.531, 0.490, 0.591],
    [0.831, 0.677, 0.758, 0.600, 0.663, 0.677, 0.611, 0.734, 0.711, 0.595, 0.657, 0.535, 0.584, 0.595, 0.544, 0.639, 0.628, 0.536, 0.585, 0.486, 0.527, 0.536, 0.494, 0.571, 0.562, 0.487, 0.528, 0.446, 0.480, 0.487, 0.452, 0.516],
    [0.677, 0.831, 0.734, 0.611, 0.677, 0.663, 0.600, 0.758, 0.595, 0.711, 0.639, 0.544, 0.595, 0.584, 0.535, 0.657, 0.536, 0.628, 0.571, 0.494, 0.536, 0.527, 0.486, 0.585, 0.487, 0.562, 0.516, 0.452, 0.487, 0.480, 0.446, 0.528],
    [0.611, 0.734, 0.831, 0.677, 0.758, 0.600, 0.663, 0.677, 0.544, 0.639, 0.711, 0.595, 0.657, 0.535, 0.584, 0.595, 0.494, 0.571, 0.628, 0.536, 0.585, 0.486, 0.527, 0.536, 0.452, 0.516, 0.562, 0.487, 0.528, 0.446, 0.480, 0.487],
    [0.600, 0.758, 0.677, 0.831, 0.734, 0.611, 0.677, 0.663, 0.535, 0.657, 0.595, 0.711, 0.639, 0.544, 0.595, 0.584, 0.486, 0.585, 0.536, 0.628, 0.571, 0.494, 0.536, 0.527, 0.446, 0.528, 0.487, 0.562, 0.516, 0.452, 0.487, 0.480],
    [0.663, 0.677, 0.611, 0.734, 0.831, 0.677, 0.758, 0.600, 0.584, 0.595, 0.544, 0.639, 0.711, 0.595, 0.657, 0.535, 0.527, 0.536, 0.494, 0.571, 0.628, 0.536, 0.585, 0.486, 0.480, 0.487, 0.452, 0.516, 0.562, 0.487, 0.528, 0.446],
    [0.677, 0.663, 0.600, 0.758, 0.677, 0.831, 0.734, 0.611, 0.595, 0.584, 0.535, 0.657, 0.595, 0.711, 0.639, 0.544, 0.536, 0.527, 0.486, 0.585, 0.536, 0.628, 0.571, 0.494, 0.487, 0.480, 0.446, 0.528, 0.487, 0.562, 0.516, 0.452],
    [0.758, 0.600, 0.663, 0.677, 0.611, 0.734, 0.831, 0.677, 0.657, 0.535, 0.584, 0.595, 0.544, 0.639, 0.711, 0.595, 0.585, 0.486, 0.527, 0.536, 0.494, 0.571, 0.628, 0.536, 0.528, 0.446, 0.480, 0.487, 0.452, 0.516, 0.562, 0.487],
    [0.734, 0.611, 0.677, 0.663, 0.600, 0.758, 0.677, 0.831, 0.639, 0.544, 0.595, 0.584, 0.535, 0.657, 0.595, 0.711, 0.571, 0.494, 0.536, 0.527, 0.486, 0.585, 0.536, 0.628, 0.516, 0.452, 0.487, 0.480, 0.446, 0.528, 0.487, 0.562],
    [0.549, 0.646, 0.720, 0.601, 0.664, 0.539, 0.590, 0.601, 0.494, 0.571, 0.628, 0.536, 0.585, 0.486, 0.527, 0.536, 0.452, 0.516, 0.562, 0.487, 0.528, 0.446, 0.480, 0.487, 0.417, 0.471, 0.509, 0.446, 0.480, 0.411, 0.440, 0.446],
    [0.646, 0.549, 0.601, 0.590, 0.539, 0.664, 0.601, 0.720, 0.571, 0.494, 0.536, 0.527, 0.486, 0.585, 0.536, 0.628, 0.516, 0.452, 0.487, 0.480, 0.446, 0.528, 0.487, 0.562, 0.471, 0.417, 0.446, 0.440, 0.411, 0.480, 0.446, 0.509],
    [0.590, 0.601, 0.549, 0.646, 0.720, 0.601, 0.664, 0.539, 0.527, 0.536, 0.494, 0.571, 0.628, 0.536, 0.585, 0.486, 0.480, 0.487, 0.452, 0.516, 0.562, 0.487, 0.528, 0.446, 0.440, 0.446, 0.417, 0.471, 0.509, 0.446, 0.480, 0.411],
    [0.601, 0.720, 0.646, 0.549, 0.601, 0.590, 0.539, 0.664, 0.536, 0.628, 0.571, 0.494, 0.536, 0.527, 0.486, 0.585, 0.487, 0.562, 0.516, 0.452, 0.487, 0.480, 0.446, 0.528, 0.446, 0.509, 0.471, 0.417, 0.446, 0.440, 0.411, 0.480],
    [0.664, 0.539, 0.590, 0.601, 0.549, 0.646, 0.720, 0.601, 0.585, 0.486, 0.527, 0.536, 0.494, 0.571, 0.628, 0.536, 0.528, 0.446, 0.480, 0.487, 0.452, 0.516, 0.562, 0.487, 0.480, 0.411, 0.440, 0.446, 0.417, 0.471, 0.509, 0.446],
    [0.539, 0.664, 0.601, 0.720, 0.646, 0.549, 0.601, 0.590, 0.486, 0.585, 0.536, 0.628, 0.571, 0.494, 0.536, 0.527, 0.446, 0.528, 0.487, 0.562, 0.516, 0.452, 0.487, 0.480, 0.411, 0.480, 0.446, 0.509, 0.471, 0.417, 0.446, 0.440],
    [0.720, 0.601, 0.664, 0.539, 0.590, 0.601, 0.549, 0.646, 0.628, 0.536, 0.585, 0.486, 0.527, 0.536, 0.494, 0.571, 0.562, 0.487, 0.528, 0.446, 0.480, 0.487, 0.452, 0.516, 0.509, 0.446, 0.480, 0.411, 0.440, 0.446, 0.417, 0.471],
    [0.601, 0.590, 0.539, 0.664, 0.601, 0.720, 0.646, 0.549, 0.536, 0.527, 0.486, 0.585, 0.536, 0.628, 0.571, 0.494, 0.487, 0.480, 0.446, 0.528, 0.487, 0.562, 0.516, 0.452, 0.446, 0.440, 0.411, 0.480, 0.446, 0.509, 0.471, 0.417],
    [0.531, 0.540, 0.498, 0.576, 0.635, 0.540, 0.591, 0.490, 0.480, 0.487, 0.452, 0.516, 0.562, 0.487, 0.528, 0.446, 0.440, 0.446, 0.417, 0.471, 0.509, 0.446, 0.480, 0.411, 0.407, 0.412, 0.387, 0.433, 0.465, 0.412, 0.441, 0.382],
    [0.540, 0.531, 0.490, 0.591, 0.540, 0.635, 0.576, 0.498, 0.487, 0.480, 0.446, 0.528, 0.487, 0.562, 0.516, 0.452, 0.446, 0.440, 0.411, 0.480, 0.446, 0.509, 0.471, 0.417, 0.412, 0.407, 0.382, 0.441, 0.412, 0.465, 0.433, 0.387],
    [0.591, 0.490, 0.531, 0.540, 0.498, 0.576, 0.635, 0.540, 0.528, 0.446, 0.480, 0.487, 0.452, 0.516, 0.562, 0.487, 0.480, 0.411, 0.440, 0.446, 0.417, 0.471, 0.509, 0.446, 0.441, 0.382, 0.407, 0.412, 0.387, 0.433, 0.465, 0.412],
    [0.576, 0.498, 0.540, 0.531, 0.490, 0.591, 0.540, 0.635, 0.516, 0.452, 0.487, 0.480, 0.446, 0.528, 0.487, 0.562, 0.471, 0.417, 0.446, 0.440, 0.411, 0.480, 0.446, 0.509, 0.433, 0.387, 0.412, 0.407, 0.382, 0.441, 0.412, 0.465],
    [0.635, 0.540, 0.591, 0.490, 0.531, 0.540, 0.498, 0.576, 0.562, 0.487, 0.528, 0.446, 0.480, 0.487, 0.452, 0.516, 0.509, 0.446, 0.480, 0.411, 0.440, 0.446, 0.417, 0.471, 0.465, 0.412, 0.441, 0.382, 0.407, 0.412, 0.387, 0.433],
    [0.540, 0.635, 0.576, 0.498, 0.540, 0.531, 0.490, 0.591, 0.487, 0.562, 0.516, 0.452, 0.487, 0.480, 0.446, 0.528, 0.446, 0.509, 0.471, 0.417, 0.446, 0.440, 0.411, 0.480, 0.412, 0.465, 0.433, 0.387, 0.412, 0.407, 0.382, 0.441],
    [0.498, 0.576, 0.635, 0.540, 0.591, 0.490, 0.531, 0.540, 0.452, 0.516, 0.562, 0.487, 0.528, 0.446, 0.480, 0.487, 0.417, 0.471, 0.509, 0.446, 0.480, 0.411, 0.440, 0.446, 0.387, 0.433, 0.465, 0.412, 0.441, 0.382, 0.407, 0.412],
    [0.490, 0.591, 0.540, 0.635, 0.576, 0.498, 0.540, 0.531, 0.446, 0.528, 0.487, 0.562, 0.516, 0.452, 0.487, 0.480, 0.411, 0.480, 0.446, 0.509, 0.471, 0.417, 0.446, 0.440, 0.382, 0.441, 0.412, 0.465, 0.433, 0.387, 0.412, 0.407],
]