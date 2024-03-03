local M = {}

---@param hay integer[]
---@param needle fun(integer): boolean
---@return integer
function M.lin_search(hay, needle)
	for i, x in ipairs(hay) do
		if needle(x) then
			return i
		end
	end
	return #hay + 1
end

---@param hay integer[]
---@param needle fun(integer): boolean
---@return integer
function M.bsearch(hay, needle)
	local lo = 1
	local hi = #hay
	while lo <= hi do
		local mid = lo + (hi - lo) // 2
		if needle(hay[mid]) then
			hi = mid - 1
		else
			lo = mid + 1
		end
	end
	return lo
end

---@param hi integer
---@param lo integer
---@param nums integer[]
---@return integer
function M.count_between(lo, hi, nums)
	local src = M.bsearch
	-- print(string.format("nums => %s", table.concat(nums, ", ")))
	-- print(string.format("lo =>%d hi=>%d", lo, hi))
	local first = src(nums, function(el)
		return el >= lo
	end)
	if first == #nums + 1 then
		return 0
	end
	local last = src(nums, function(el)
		return el > hi
	end) - 1
	local count = last - first + 1
	-- print(string.format("count =>%d", count))
	-- print("=======")
	return count
end

return M
