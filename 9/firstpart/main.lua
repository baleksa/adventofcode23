local function all_zeros(tbl)
	for _, v in ipairs(tbl) do
		if v ~= 0 then
			return false
		end
	end
	return true
end
local function next_value(history)
	local seqs = { history }
	local seq = seqs[1]
	while not all_zeros(seq) do
		local new_seq = {}
		for i = 2, #seq do
			new_seq[#new_seq + 1] = seq[i] - seq[i - 1]
		end
		seq = new_seq
		seqs[#seqs + 1] = seq
	end
	for _, s in ipairs(seqs) do
		print(table.concat(s, ", "))
	end
	local result = 0
	for i = #seqs - 1, 1, -1 do
		result = result + seqs[i][#seqs[i]]
	end
	print(result)
	return result
end

local sum = 0

for line in io.lines(arg[1]) do
	print(line)
	local history = {}
	for x in string.gmatch(line, "-?%d+") do
		history[#history + 1] = x
	end
	print(table.concat(history, ", "))
	sum = sum + next_value(history)
end

print(sum)
