<div align="center">

# Ahoy
##### Easy docker management in web browser.


<img src="https://github.com/radim91/ahoy/assets/34510551/6b458641-38e7-49d8-a3e5-cbfec9ef492a" width="100" />
</div>

## Quick How to
1. Download latest binary (only Linux).
2. Run the binary.
3. Visit http://localhost:8080 in your web browser.
   
<br/><br/>
![Snímek obrazovky z 2024-05-19 14-34-50](https://github.com/radim91/ahoy/assets/34510551/9a5be7d4-d4fc-47ed-901c-dfd176171162)

## Additional How To
By default Ahoy runs on port 8080 on localhost. But you can configure application URL. Just create `.env` file in the directory of Ahoy's binary like this:
```
URL="localhost:1337"
```
If neccessary, don't forget to set the binary as executable:
```
chmod +x /PATH/TO/BINARY
```

