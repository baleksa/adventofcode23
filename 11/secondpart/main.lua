#!/usr/bin/env lua5.4

---@alias Galaxy {x: integer, y: integer}

---@param line string
---@return boolean
local function is_empty_line(line)
	for c in string.gmatch(line, ".") do
		if c ~= "." then
			return false
		end
	end
	return true
end

---@param g1 Galaxy
---@param g2 Galaxy
---@param er integer[]
---@param ec integer[]
---@param space_multiplier integer
---@return integer
local function man_dist(g1, g2, er, ec, space_multiplier)
	local x1 = g1.x
	local y1 = g1.y
	local x2 = g2.x
	local y2 = g2.y
	local difx = math.abs(x1 - x2)
	local dify = math.abs(y1 - y2)
	local dist = difx + dify

	local n_empty_rows_between = require("bsearch").count_between(math.min(x1, x2), math.max(x1, x2), er)
	dist = dist + n_empty_rows_between * (space_multiplier - 1)

	local n_empty_clmns_between = require("bsearch").count_between(math.min(y1, y2), math.max(y1, y2), ec)
	dist = dist + n_empty_clmns_between * (space_multiplier - 1)

	return dist
end

---@param galaxies Galaxy[]
---@param empty_rows integer[]
---@param empty_columns integer[]
---@param space_multiplier integer
---@return integer
local function sum_man_dist(galaxies, empty_rows, empty_columns, space_multiplier)
	local sum = 0
	for i = 1, #galaxies do
		for j = i + 1, #galaxies do
			sum = sum + man_dist(galaxies[i], galaxies[j], empty_rows, empty_columns, space_multiplier)
		end
	end
	print(string.format("Multiplier =>%15d Sum =>%15d", space_multiplier, sum))
	return sum
end

---@param j integer
---@param image string[]
---@return boolean
local function is_empty_column(j, image)
	for i = 1, #image do
		if string.sub(image[i], j, j) ~= "." then
			return false
		end
	end
	return true
end

---@type string[]
local image = {}
---@type integer[]
local empty_rows = {}
---@type integer[]
local empty_columns = {}

---@type Galaxy[]
local galaxies = {}

local line_num = 1
for line in io.lines(arg[1]) do
	image[#image + 1] = line
	if is_empty_line(line) then
		empty_rows[#empty_rows + 1] = line_num
	end
	line_num = line_num + 1
end

-- for _, line in pairs(image) do
-- 	print(line)
-- end

print(string.format("Empty rows => %s", table.concat(empty_rows, ", ")))

for j = 1, #image[1] do
	if is_empty_column(j, image) then
		empty_columns[#empty_columns + 1] = j
	end
end
print(string.format("Empty columns => %s", table.concat(empty_columns, ", ")))

for i = 1, #image do
	for j = 1, #image[i] do
		if string.sub(image[i], j, j) == "#" then
			galaxies[#galaxies + 1] = { x = i, y = j }
			-- print(string.format("Galaxy => {%4d,%4d}", i, j))
		end
	end
end

local muls = { 2, 10, 100, 1000000 }
for _, mul in pairs(muls) do
	sum_man_dist(galaxies, empty_rows, empty_columns, mul)
end
