
local addonName = "TheLedger"
local addonDB

function SaveToFile(filename, dataString)
    local file = assert(io.open(filename, "w"))
    file:write(dataString)
    file:close()
end


print("This is the Start")
local frame = CreateFrame("Frame")

frame:RegisterEvent("PLAYER_LOGIN")

-- frame:RegisterEvent("AUCTION_HOUSE_SHOW")
-- frame:RegisterEvent("AUCTION_CANCELED")
-- frame:RegisterEvent("AUCTION_HOUSE_AUCTION_CREATED")
-- frame:RegisterEvent("AUCTION_HOUSE_AUCTIONS_EXPIRED")
-- frame:RegisterEvent("AUCTION_HOUSE_BROWSE_RESULTS_ADDED")
-- frame:RegisterEvent("AUCTION_HOUSE_BROWSE_RESULTS_UPDATED")
-- frame:RegisterEvent("AUCTION_HOUSE_CLOSED")
-- frame:RegisterEvent("AUCTION_HOUSE_ITEM_DELIVERY_DELAY_UPDATE")
-- frame:RegisterEvent("AUCTION_HOUSE_NEW_BID_RECEIVED")
-- frame:RegisterEvent("AUCTION_HOUSE_NEW_RESULTS_RECEIVED")
-- frame:RegisterEvent("AUCTION_HOUSE_PURCHASE_COMPLETED")
-- frame:RegisterEvent("AUCTION_HOUSE_SHOW_COMMODITY_WON_NOTIFICATION")
-- frame:RegisterEvent("AUCTION_HOUSE_SHOW_NOTIFICATION")
-- frame:RegisterEvent("AUCTION_HOUSE_THROTTLED_SYSTEM_READY")
-- frame:RegisterEvent("AUCTION_HOUSE_THROTTLED_MESSAGE_QUEUED")
-- frame:RegisterEvent("AUCTION_HOUSE_THROTTLED_MESSAGE_DROPPED")
-- frame:RegisterEvent("AUCTION_HOUSE_THROTTLED_MESSAGE_SENT")
-- frame:RegisterEvent("AUCTION_MULTISELL_START")
-- frame:RegisterEvent("AUCTION_MULTISELL_UPDATE")
-- frame:RegisterEvent("COMMODITY_PRICE_UPDATED")
-- frame:RegisterEvent("COMMODITY_PURCHASE_SUCCEEDED")
-- frame:RegisterEvent("COMMODITY_PURCHASED")
-- frame:RegisterEvent("COMMODITY_SEARCH_RESULTS_ADDED")
-- frame:RegisterEvent("COMMODITY_SEARCH_RESULTS_UPDATED")
-- frame:RegisterEvent("ITEM_KEY_ITEM_INFO_RECEIVED")
-- frame:RegisterEvent("ITEM_PURCHASED")
-- frame:RegisterEvent("ITEM_SEARCH_RESULTS_ADDED")
-- frame:RegisterEvent("ITEM_SEARCH_RESULTS_UPDATED")
-- frame:RegisterEvent("OWNED_AUCTIONS_UPDATED")
-- frame:RegisterEvent("REPLICATE_ITEM_LIST_UPDATE")

frame:SetScript("OnEvent", function(self, event, ...)
    local p1, p2 = ...
    local currentTime = date("%Y-%m-%d %H:%M:%S")
    print(currentTime, " Event Captured", event, PrintTable(p1, 0), p2)
    if event == "PLAYER_LOGIN" then
        print("Before")
        -- local res = ReadAndPrintBagContents()

        print("Welcome to World of Warcraft! Toggle")
    end

    if event == "AUCTION_HOUSE_SHOW" then
        C_Timer.After(10, function()
            local items = {
                { itemID = 190324}, -- awakened order
                { itemID = 191461}, -- hochenblume 2
                { itemID = 191462}, -- hochenblume 3
                { itemID = 191464}, -- saxifage 1
                { itemID = 191359}, -- phial of elemental chaos 3
                { itemID = 191341}, -- phial of tepid versatility 3
            }
            local sorts = {
                {sortOrder = Enum.AuctionHouseSortOrder.Price, reverseSort = false},
            }
            C_AuctionHouse.SearchForItemKeys(items, sorts)
        end)
        print("auction house show")

        
    end

    if event == "AUCTION_HOUSE_NEW_RESULTS_RECEIVED" or event == "AUCTION_HOUSE_BROWSE_RESULTS_UPDATED" then
        local browseResults = C_AuctionHouse.GetBrowseResults()
        for i, item in ipairs(browseResults) do
            PrintTable(item, 0)
        end
    end

end)

-- /console scriptErrors 1  enables error reporting
-- /console scriptErrors 0

-- save variables to file https://wowwiki-archive.fandom.com/wiki/Saving_variables_between_game_sessions#:~:text=Variables%20that%20are%20saved%20and,toc%20file.

-- 38b0218c1fb24e4599b60c358b01964d
-- tMzdyGzfOE01bCzc6yccCd186MK26TWG


