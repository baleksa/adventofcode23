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

---@param x1 integer
---@param y1 integer
---@param x2 integer
---@param y2 integer
---@return integer
local function man_dist(x1, y1, x2, y2)
	local difx = x1 - x2
	local dify = y1 - y2
	if difx < 0 then
		difx = -difx
	end
	if dify < 0 then
		dify = -dify
	end

	return difx + dify
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

local function print_table(tbl)
	print(table.concat(tbl, ", "))
end

local image = {}
for line in io.lines(arg[1]) do
	image[#image + 1] = line
	if is_empty_line(line) then
		image[#image + 1] = string.sub(line, 1)
	end
end

for _, line in pairs(image) do
	print(line)
end

local empty_colums = {}
for j = 1, #image[1] do
	if is_empty_column(j, image) then
		empty_colums[#empty_colums + 1] = j
	end
end
print_table(empty_colums)

for j = #empty_colums, 1, -1 do
	for k, line in ipairs(image) do
		image[k] = string.sub(line, 1, empty_colums[j]) .. "." .. string.sub(line, empty_colums[j] + 1)
	end
end

for _, line in pairs(image) do
	print(line)
end

local galaxies = {}
for i = 1, #image do
	for j = 1, #image[i] do
		if string.sub(image[i], j, j) == "#" then
			galaxies[#galaxies + 1] = { x = i, y = j }
			print(i, j)
		end
	end
end

local sum = 0
for i = 1, #galaxies do
	for j = i + 1, #galaxies do
		sum = sum + man_dist(galaxies[i].x, galaxies[i].y, galaxies[j].x, galaxies[j].y)
	end
end

print("Sum =>", sum)
