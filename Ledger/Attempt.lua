
function OuputVarContents(var)
    for key,value in pairs(item) do
        print("found member " .. key);
    end
end

function PrintTable(t, indent)
    if (t == nil) then
        return;
    end
    indent = indent or 0

    if type(t) ~= "table" then
        print(string.rep("  ", indent) .. t)
        return;
    end
    local keys = {}

    for k, v in pairs(t) do
        table.insert(keys, k)
    end

    table.sort(keys)

    for _, k in ipairs(keys) do
        local v = t[k]
        local keyStr = tostring(k)

        if type(v) == "table" then
            print(string.rep("  ", indent) .. keyStr .. " = {")
            PrintTable(v, indent + 1)
            print(string.rep("  ", indent) .. "}")
        else
            print(string.rep("  ", indent) .. keyStr .. " = " .. tostring(v))
        end
    end
end


function ReadAndPrintBagContents()
    itemData = {}
    for containerIndex = Enum.BagIndex.Backpack, Constants.InventoryConstants.NumBagSlots do
        local slots = C_Container.GetContainerNumSlots(containerIndex);
        print("container", containerIndex, "slots", slots)
        if (slots > 0) then
            for slotIndex = 1, slots do
                local itemInfo = C_Container.GetContainerItemInfo(containerIndex, slotIndex);
                if (itemInfo ~= nil) then
                    local item = Item:CreateFromItemID(itemInfo.itemID);
                    -- local loc = ItemLocation:CreateFormBagAndSlot(containerIndex, slotIndex)
                    -- print(containerIndex, slotIndex, itemInfo.itemID, item:GetItemName(), itemInfo.stackCount, itemInfo.quality)
                    itemData[item.itemID] = itemInfo.stackCount
                end
            end
        end
    end
    return itemData
end